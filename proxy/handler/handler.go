package handler

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"text/template"
	"time"

	"rest-to-soap/proxy/config"
	parser "rest-to-soap/proxy/generated"
	transport "rest-to-soap/proxy/soap"
	"rest-to-soap/proxy/wsdl"

	"go.uber.org/zap"
)

// RequestBody represents the XML structure for SOAP requests
type RequestBody struct {
	XMLName xml.Name    `xml:"request"`
	Data    interface{} `xml:",any"`
}

// Handler handles HTTP requests and forwards them to SOAP endpoints
type Handler struct {
	cfg          *config.Config
	client       *transport.Client
	pool         *Pool
	logger       *zap.Logger
	wsdl         *wsdl.Parser
	requestTmpl  *template.Template
	responseTmpl *template.Template
}

// NewHandler creates a new request handler
func NewHandler(cfg *config.Config, logger *zap.Logger) (*Handler, error) {
	// Load templates
	requestTmpl, err := template.ParseFiles("templates/request.tmpl")
	if err != nil {
		return nil, err
	}

	responseTmpl, err := template.ParseFiles("templates/response.tmpl")
	if err != nil {
		return nil, err
	}

	return &Handler{
		cfg:          cfg,
		client:       transport.NewClient(30*time.Second, logger),
		pool:         NewPool(),
		logger:       logger,
		wsdl:         wsdl.NewParser(logger),
		requestTmpl:  requestTmpl,
		responseTmpl: responseTmpl,
	}, nil
}

// ServeHTTP implements http.Handler
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Find route
	route := h.findRoute(path)
	if route == nil {
		http.NotFound(w, r)
		return
	}

	// Parse request body if present
	var body map[string]interface{}
	if r.Body != nil && r.ContentLength > 0 {
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			h.logger.Error("Failed to parse request body", zap.Error(err))
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
	}

	// Process request in worker pool
	err := h.pool.WithContext(r.Context(), func() error {
		return h.processRequest(w, r, route, body)
	})

	if err != nil {
		h.logger.Error("Request processing failed", zap.Error(err))
		// Return error as JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
}

func (h *Handler) findRoute(path string) *config.RouteConfig {
	for _, route := range h.cfg.Routes {
		if route.Path == path {
			return &route
		}
	}
	return nil
}

func (h *Handler) processRequest(w http.ResponseWriter, r *http.Request, route *config.RouteConfig, body map[string]interface{}) error {
	// Get WSDL type info if available
	var typeInfo map[string]interface{}
	if route.WSDLURL != "" {
		var err error
		typeInfo, err = h.wsdl.GetTypeInfo(route.WSDLURL)
		if err != nil {
			h.logger.Warn("Failed to get WSDL type info", zap.Error(err))
		}
	}

	// Convert body to XML with type info if available
	var bodyXML []byte
	var err error
	if body != nil {
		if typeInfo != nil {
			bodyXML, err = xml.MarshalIndent(body, "", "  ")
		} else {
			reqBody := RequestBody{Data: body}
			bodyXML, err = xml.Marshal(reqBody)
		}
		if err != nil {
			return err
		}
	}

	// Render SOAP request
	var soapReq bytes.Buffer
	if err := h.requestTmpl.Execute(&soapReq, map[string]interface{}{
		"Action": route.Headers["SOAPAction"],
		"Body":   string(bodyXML),
	}); err != nil {
		return err
	}

	// Log the SOAP request
	h.logger.Info("Sending SOAP request",
		zap.String("endpoint", route.SoapEndpoint),
		zap.String("action", route.Headers["SOAPAction"]),
		zap.String("request", fmt.Sprintf("%q", soapReq.String())),
	)

	// Create SOAP request
	req, err := http.NewRequestWithContext(r.Context(), "POST", route.SoapEndpoint, &soapReq)
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")
	for k, v := range route.Headers {
		req.Header.Set(k, v)
	}

	// Log headers
	h.logger.Info("Request headers",
		zap.Any("headers", req.Header),
	)

	// Send request
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		// Try to parse as SOAP fault first
		var soapFault struct {
			XMLName xml.Name `xml:"Envelope"`
			Body    struct {
				Fault *struct {
					FaultCode   string `xml:"faultcode"`
					FaultString string `xml:"faultstring"`
				} `xml:"Fault"`
			} `xml:"Body"`
		}

		if err := xml.Unmarshal(respBody, &soapFault); err == nil && soapFault.Body.Fault != nil {
			return &SoapFault{
				Code:   soapFault.Body.Fault.FaultCode,
				String: soapFault.Body.Fault.FaultString,
			}
		}

		// If not a SOAP fault, return a generic error with the response body
		return fmt.Errorf("SOAP service returned error (status %d): %s", resp.StatusCode, string(respBody))
	}

	// Parse SOAP response using the appropriate parser
	response, err := parser.Parse(respBody)
	if err != nil {
		return fmt.Errorf("failed to parse SOAP response: %w", err)
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(response))
	return err
}

// SoapFault represents a SOAP fault response
type SoapFault struct {
	Code   string
	String string
}

func (f *SoapFault) Error() string {
	return f.String
}
