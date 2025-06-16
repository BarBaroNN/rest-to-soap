package config

import (
	"encoding/json"
	"os"
	"time"
)

// Config represents the application configuration
type Config struct {
	Server  ServerConfig  `json:"server"`
	Routes  []RouteConfig `json:"routes"`
	Logging LogConfig     `json:"logging"`
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Port         int           `json:"port"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	IdleTimeout  time.Duration `json:"idle_timeout"`
}

// LogConfig defines logging configuration
type LogConfig struct {
	Level  string `json:"level"`
	Format string `json:"format"`
}

// RouteConfig represents a route configuration
type RouteConfig struct {
	Path             string            `json:"path"`
	Method           string            `json:"method"`
	SoapEndpoint     string            `json:"soap_endpoint"`
	RequestTemplate  string            `json:"request_template"`
	ResponseTemplate string            `json:"response_template"`
	Headers          map[string]string `json:"headers"`
	WSDLURL          string            `json:"wsdl_url,omitempty"`
	Timeout          time.Duration     `json:"timeout"`
}

// Load loads the configuration from a file
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// UnmarshalJSON implements custom JSON unmarshaling for time.Duration fields
func (s *ServerConfig) UnmarshalJSON(data []byte) error {
	type Alias ServerConfig
	aux := &struct {
		ReadTimeout  string `json:"read_timeout"`
		WriteTimeout string `json:"write_timeout"`
		IdleTimeout  string `json:"idle_timeout"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var err error
	s.ReadTimeout, err = time.ParseDuration(aux.ReadTimeout)
	if err != nil {
		return err
	}

	s.WriteTimeout, err = time.ParseDuration(aux.WriteTimeout)
	if err != nil {
		return err
	}

	s.IdleTimeout, err = time.ParseDuration(aux.IdleTimeout)
	if err != nil {
		return err
	}

	return nil
}

// UnmarshalJSON implements custom JSON unmarshaling for RouteConfig
func (r *RouteConfig) UnmarshalJSON(data []byte) error {
	type Alias RouteConfig
	aux := &struct {
		Timeout string `json:"timeout"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var err error
	r.Timeout, err = time.ParseDuration(aux.Timeout)
	if err != nil {
		return err
	}

	return nil
}
