package handler

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"html/template"
	"io"
	"net/http"
	"time"

	"rest-to-soap/proxy/config"
	transport "rest-to-soap/proxy/soap"
	"rest-to-soap/proxy/wsdl"

	"go.uber.org/zap"
)

// RequestBody represents the XML structure for SOAP requests
type RequestBody struct {
	XMLName xml.Name    `xml:"request"`
	Data    interface{} `xml:",any"`
}

// Handler processes HTTP requests and forwards them to SOAP endpoints
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

	// Parse request body
	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		h.logger.Error("Failed to parse request body", zap.Error(err))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Process request in worker pool
	err := h.pool.WithContext(r.Context(), func() error {
		return h.processRequest(w, r, route, body)
	})

	if err != nil {
		h.logger.Error("Request processing failed", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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
	if typeInfo != nil {
		bodyXML, err = xml.MarshalIndent(body, "", "  ")
	} else {
		reqBody := RequestBody{Data: body}
		bodyXML, err = xml.Marshal(reqBody)
	}
	if err != nil {
		return err
	}

	// Render SOAP request
	var soapReq bytes.Buffer
	if err := h.requestTmpl.Execute(&soapReq, map[string]interface{}{
		"Action": route.Headers["SOAPAction"],
		"Body":   string(bodyXML),
	}); err != nil {
		return err
	}

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

	// Parse SOAP response
	var soapResp struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			Content []byte `xml:",innerxml"`
		} `xml:"Body"`
		Fault *struct {
			FaultCode   string `xml:"faultcode"`
			FaultString string `xml:"faultstring"`
		} `xml:"Fault"`
	}

	if err := xml.Unmarshal(respBody, &soapResp); err != nil {
		return err
	}

	// Check for SOAP fault
	if soapResp.Fault != nil {
		return &SoapFault{
			Code:   soapResp.Fault.FaultCode,
			String: soapResp.Fault.FaultString,
		}

	}

	// Convert to JSON
	var jsonResp map[string]interface{}
	if err := xml.Unmarshal(soapResp.Body.Content, &jsonResp); err != nil {
		return err
	}

	// Render JSON response
	w.Header().Set("Content-Type", "application/json")
	return h.responseTmpl.Execute(w, map[string]interface{}{
		"Status": "success",
		"Data":   jsonResp,
	})
}

// SoapFault represents a SOAP fault response
type SoapFault struct {
	Code   string
	String string
}

func (f *SoapFault) Error() string {
	return f.String
}
