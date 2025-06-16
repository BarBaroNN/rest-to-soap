package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rest-to-soap/config"
	"go.uber.org/zap"
)

func TestHandler(t *testing.T) {
	// Create test configuration
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port:         8080,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		Routes: []config.RouteConfig{
			{
				Path:             "/test",
				Method:           "POST",
				SoapEndpoint:     "http://example.com/soap",
				RequestTemplate:  "templates/request.tmpl",
				ResponseTemplate: "templates/response.tmpl",
				Headers: map[string]string{
					"SOAPAction": "http://example.com/action",
				},
				Timeout: 30 * time.Second,
			},
		},
		Logging: config.LogConfig{
			Level:  "debug",
			Format: "console",
		},
	}

	// Create logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	// Create handler
	handler, err := NewHandler(cfg, logger)
	if err != nil {
		t.Fatalf("Failed to create handler: %v", err)
	}

	// Create test request
	reqBody := map[string]interface{}{
		"test": "value",
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest("POST", "/test", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Create response recorder
	w := httptest.NewRecorder()

	// Serve request
	handler.ServeHTTP(w, req)

	// Check response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Check response body
	var respBody map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&respBody); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	// Verify response structure
	if status, ok := respBody["status"].(string); !ok || status != "success" {
		t.Errorf("Expected status 'success', got %v", status)
	}

	if data, ok := respBody["data"].(map[string]interface{}); !ok {
		t.Error("Expected 'data' field in response")
	} else if testValue, ok := data["test"].(string); !ok || testValue != "value" {
		t.Errorf("Expected test value 'value', got %v", testValue)
	}
}

func TestHandlerNotFound(t *testing.T) {
	// Create test configuration
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port:         8080,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		Routes: []config.RouteConfig{},
		Logging: config.LogConfig{
			Level:  "debug",
			Format: "console",
		},
	}

	// Create logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	// Create handler
	handler, err := NewHandler(cfg, logger)
	if err != nil {
		t.Fatalf("Failed to create handler: %v", err)
	}

	// Create test request
	req := httptest.NewRequest("POST", "/not-found", nil)
	w := httptest.NewRecorder()

	// Serve request
	handler.ServeHTTP(w, req)

	// Check response
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}

	// Check response body
	var respBody map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&respBody); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	// Verify error response structure
	if status, ok := respBody["status"].(string); !ok || status != "error" {
		t.Errorf("Expected status 'error', got %v", status)
	}

	if errObj, ok := respBody["error"].(map[string]interface{}); !ok {
		t.Error("Expected 'error' field in response")
	} else if message, ok := errObj["message"].(string); !ok || message != "Not Found" {
		t.Errorf("Expected error message 'Not Found', got %v", message)
	}
}

func TestHandlerInvalidBody(t *testing.T) {
	// Create test configuration
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port:         8080,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		Routes: []config.RouteConfig{
			{
				Path:             "/test",
				Method:           "POST",
				SoapEndpoint:     "http://example.com/soap",
				RequestTemplate:  "templates/request.tmpl",
				ResponseTemplate: "templates/response.tmpl",
				Timeout:          30 * time.Second,
			},
		},
		Logging: config.LogConfig{
			Level:  "debug",
			Format: "console",
		},
	}

	// Create logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	// Create handler
	handler, err := NewHandler(cfg, logger)
	if err != nil {
		t.Fatalf("Failed to create handler: %v", err)
	}

	// Create test request with invalid JSON
	req := httptest.NewRequest("POST", "/test", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Serve request
	handler.ServeHTTP(w, req)

	// Check response
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Check response body
	var respBody map[string]interface{}
	if err := json.NewDecoder(w.Body).Decode(&respBody); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	// Verify error response structure
	if status, ok := respBody["status"].(string); !ok || status != "error" {
		t.Errorf("Expected status 'error', got %v", status)
	}

	if errObj, ok := respBody["error"].(map[string]interface{}); !ok {
		t.Error("Expected 'error' field in response")
	} else if message, ok := errObj["message"].(string); !ok || message == "" {
		t.Error("Expected non-empty error message")
	}
}
