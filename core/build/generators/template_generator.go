package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"rest-to-soap/core/config"
)

// TemplateGenerator handles the generation of Go templates from WSDL files
type TemplateGenerator struct {
	outputDir string
}

// NewTemplateGenerator creates a new template generator
func NewTemplateGenerator() *TemplateGenerator {
	// Create the output directory if it doesn't exist
	outputDir := "pkg/generated"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create output directory: %v", err))
	}

	return &TemplateGenerator{
		outputDir: outputDir,
	}
}

// GenerateTemplates generates Go templates for all routes in the configuration
func (g *TemplateGenerator) GenerateTemplates(cfg *config.Config) error {
	for _, route := range cfg.Routes {
		if route.WSDLURL == "" {
			continue // Skip routes without WSDL
		}

		// Use SOAPAction as the operation name
		operationName := route.SoapAction
		if operationName == "" {
			continue // Skip if no SOAPAction is defined
		}

		// Generate the template
		if err := g.generateTemplate(route.WSDLURL, operationName, route.ResponseTemplate); err != nil {
			return fmt.Errorf("failed to generate template for operation %s: %w", operationName, err)
		}
	}
	return nil
}

// generateTemplate generates a Go template for a specific WSDL operation
func (g *TemplateGenerator) generateTemplate(wsdlURL, operationName, templatePath string) error {
	// Generate structs from WSDL
	structs, responseType, err := ExtractStructsFromWSDL(wsdlURL, operationName)
	if err != nil {
		return fmt.Errorf("failed to extract structs: %w", err)
	}

	// Create the output file
	outputPath := filepath.Join(g.outputDir, fmt.Sprintf("%s_parser.go", operationName))

	// Prepare the generated code
	generatedCode := fmt.Sprintf(
		`package generated

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

%s

// %sParser parses the SOAP response for the %s operation
func %sParse(xmlData []byte) (string, error) {
	fmt.Println(string(xmlData))

	// Define the SOAP envelope structure with the proper response type
	var response struct {
		XMLName xml.Name %s
		Body    struct {
			Response %v %v
		} %s
	}

	// Unmarshal the XML into our strongly-typed struct
	if err := xml.Unmarshal(xmlData, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal XML: %%w", err)
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("%s")
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %%w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, response.Body.Response); err != nil {
		return "", fmt.Errorf("failed to execute template: %%w", err)
	}
	return buf.String(), nil
}
`,
		structs,
		operationName,
		operationName,
		operationName,
		"`xml:\"http://schemas.xmlsoap.org/soap/envelope/ Envelope\"`",
		responseType,
		fmt.Sprintf("`xml:\"%s\"`", responseType),
		"`xml:\"http://schemas.xmlsoap.org/soap/envelope/ Body\"`",
		templatePath,
	)

	// Create or update the file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	if _, err := outputFile.WriteString(generatedCode); err != nil {
		return fmt.Errorf("failed to write generated code: %w", err)
	}

	return nil
}
