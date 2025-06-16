package test

import (
	"path/filepath"
	"runtime"
	"testing"

	"rest-to-soap/pkg/core"
	"rest-to-soap/pkg/validation"
)

// GetTestFixturesDir returns the path to the test fixtures directory
func GetTestFixturesDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "fixtures")
}

// CreateTestRegistry creates a registry with test configuration
func CreateTestRegistry(t *testing.T) (*core.DefaultRegistry, *validation.Validator) {
	parser := core.NewParser()
	cache := core.NewMemoryCache()
	registry := core.NewRegistry(parser, cache)
	validator := validation.NewValidator()

	// Add WSDL paths for validation
	validator.AddWSDLPaths("example", []string{
		"ISOCode",
		"Name",
		"CapitalCity",
		"Languages",
		"Languages.Name",
		"Languages.Code",
		"Languages.IsOfficial",
		"Languages.Speakers",
		"Population",
		"Area",
		"Currency",
		"Currency.Code",
		"Currency.Name",
		"Currency.Symbol",
		"Currency.ExchangeRate",
		"Currency.IsActive",
		"Timezones",
		"Borders",
		"Region",
		"Subregion",
		"Flag",
		"IsIndependent",
		"GDP",
		"Gini",
		"HDI",
		"Coordinates",
		"Coordinates.Latitude",
		"Coordinates.Longitude",
		"Coordinates.IsValid",
	})

	return registry, validator
}

// LoadTestTemplate loads a template from the test fixtures
func LoadTestTemplate(t *testing.T, registry *core.DefaultRegistry, name string) {
	path := filepath.Join(GetTestFixturesDir(), name+".tmpl")
	if err := registry.LoadTemplate(name, path); err != nil {
		t.Fatalf("Failed to load test template: %v", err)
	}
}

// CreateTestData creates sample data for testing
func CreateTestData() map[string]interface{} {
	return map[string]interface{}{
		"ISOCode":     "US",
		"Name":        "United States",
		"CapitalCity": "Washington, D.C.",
		"Languages": []map[string]interface{}{
			{
				"Name":       "English",
				"Code":       "en",
				"IsOfficial": true,
				"Speakers":   231000000,
			},
			{
				"Name":       "Spanish",
				"Code":       "es",
				"IsOfficial": false,
				"Speakers":   41000000,
			},
		},
		"Population": 331002651,
		"Area":       9833517,
		"Currency": map[string]interface{}{
			"Code":         "USD",
			"Name":         "United States Dollar",
			"Symbol":       "$",
			"ExchangeRate": 1.0,
			"IsActive":     true,
		},
		"Timezones": []string{
			"UTC-12:00",
			"UTC-11:00",
			"UTC-10:00",
		},
		"Borders": []string{
			"CAN",
			"MEX",
		},
		"Region":        "Americas",
		"Subregion":     "North America",
		"Flag":          "https://flagcdn.com/us.svg",
		"IsIndependent": true,
		"GDP":           20.95e12, // 20.95 trillion
		"Gini":          41.4,
		"HDI":           0.921,
		"Coordinates": map[string]interface{}{
			"Latitude":  37.0902,
			"Longitude": -95.7129,
			"IsValid":   true,
		},
	}
}
