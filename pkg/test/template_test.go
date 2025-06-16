package test

import (
	"fmt"
	"strings"
	"testing"

	"rest-to-soap/pkg/validation"
)

func TestTemplateSystem(t *testing.T) {
	// Create test registry and validator
	registry, validator := CreateTestRegistry(t)

	// Load test template
	LoadTestTemplate(t, registry, "example")

	// Get template for validation
	tmpl, err := registry.GetTemplate("example")
	if err != nil {
		t.Fatalf("Failed to get template: %v", err)
	}

	// Validate template
	errors, err := validator.ValidateTemplate("example", tmpl)
	if err != nil {
		t.Fatalf("Failed to validate template: %v", err)
	}

	// Check for validation errors
	if len(errors) > 0 {
		t.Errorf("Template validation errors:\n%s", validation.FormatValidationErrors(errors))
	}

	// Test template execution
	data := CreateTestData()
	result, err := registry.ExecuteTemplate("example", data)
	if err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	// Verify the result contains expected content
	expectedFields := []string{
		`"ISOCode": "US"`,
		`"name": "United States"`,
		`"capitalCity": "Washington, D.C."`,
		`"languages": [`,
		`"isOfficial": true`,
		`"speakers": 231000000`,
		`"population": 331002651`,
		`"area": 9833517`,
		`"currency": {`,
		`"exchangeRate": 1`,
		`"isActive": true`,
		`"timezones": [`,
		`"borders": [`,
		`"region": "Americas"`,
		`"subregion": "North America"`,
		`"flag": "https://flagcdn.com/us.svg"`,
		`"isIndependent": true`,
		`"gdp": 2.095e+13`,
		`"gini": 41.4`,
		`"hdi": 0.921`,
		`"coordinates": {`,
		`"latitude": 37.0902`,
		`"longitude": -95.7129`,
		`"isValid": true`,
	}

	missing := false
	for _, field := range expectedFields {
		if !strings.Contains(result, field) {
			t.Errorf("Expected field not found in result: %s", field)
			missing = true
		}
	}
	if missing {
		fmt.Println("\n--- TEMPLATE OUTPUT START ---\n" + result + "\n--- TEMPLATE OUTPUT END ---\n")
	}
}

func TestTemplateValidation(t *testing.T) {
	// Create test registry and validator
	registry, validator := CreateTestRegistry(t)

	// Load test template
	LoadTestTemplate(t, registry, "example")

	// Get WSDL paths
	paths, err := validator.GetWSDLPaths("example")
	if err != nil {
		t.Fatalf("Failed to get WSDL paths: %v", err)
	}

	// Test validation with missing required path
	validator.AddWSDLPaths("example", append(paths, "NewField"))
	tmpl, err := registry.GetTemplate("example")
	if err != nil {
		t.Fatalf("Failed to get template: %v", err)
	}

	errors, err := validator.ValidateTemplate("example", tmpl)
	if err != nil {
		t.Fatalf("Failed to validate template: %v", err)
	}

	if len(errors) == 0 {
		t.Error("Expected validation errors for missing required path")
	}

	// Test validation with extra path
	tmpl.Paths = append(tmpl.Paths, "ExtraField")
	errors, err = validator.ValidateTemplate("example", tmpl)
	if err != nil {
		t.Fatalf("Failed to validate template: %v", err)
	}

	if len(errors) == 0 {
		t.Error("Expected validation errors for extra path")
	}
}

func TestTemplateCache(t *testing.T) {
	// Create test registry
	registry, _ := CreateTestRegistry(t)

	// Load test template
	LoadTestTemplate(t, registry, "example")

	// Test template retrieval
	tmpl, err := registry.GetTemplate("example")
	if err != nil {
		t.Fatalf("Failed to get template: %v", err)
	}

	if tmpl == nil {
		t.Error("Expected template to be found in cache")
	}

	// Test template listing
	templates := registry.ListTemplates()
	if len(templates) != 1 || templates[0] != "example" {
		t.Errorf("Expected template list to contain 'example', got %v", templates)
	}
}
