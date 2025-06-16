package genparser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"rest-to-soap/proxy/config"
)

// TemplateGenerator handles the generation of Go templates from WSDL files
type TemplateGenerator struct {
	outputDir string
}

// NewTemplateGenerator creates a new template generator
func NewTemplateGenerator() *TemplateGenerator {
	// Create the output directory if it doesn't exist
	outputDir := "/etc/generated-template"
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
func (g *TemplateGenerator) generateTemplate(wsdlURL, operationName, templateContent string) error {
	// Generate structs from WSDL
	structs, err := ExtractStructsFromWSDL(wsdlURL, operationName)
	if err != nil {
		return fmt.Errorf("failed to extract structs: %w", err)
	}

	// Create the output file
	outputPath := filepath.Join(g.outputDir, fmt.Sprintf("%s.go", operationName))
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// Write the package declaration
	if _, err := fmt.Fprintf(outputFile, "package generated\n\n"); err != nil {
		return fmt.Errorf("failed to write package declaration: %w", err)
	}

	// Write the imports
	if _, err := fmt.Fprintf(outputFile, `import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

`); err != nil {
		return fmt.Errorf("failed to write imports: %w", err)
	}

	// Write the generated structs
	if _, err := fmt.Fprintf(outputFile, "%s\n\n", structs); err != nil {
		return fmt.Errorf("failed to write structs: %w", err)
	}

	// Write the parser function
	if _, err := fmt.Fprintf(outputFile, `// Parser parses the SOAP response and renders JSON using the template
func Parser(xmlData []byte) (string, error) {
	// Define the SOAP envelope structure with the proper response type
	var response struct {
		XMLName xml.Name `+"`xml:\"http://schemas.xmlsoap.org/soap/envelope/ Envelope\"`"+`
		Body    struct {
			Response %sResponse `+"`xml:\"Response\"`"+`
		} `+"`xml:\"http://schemas.xmlsoap.org/soap/envelope/ Body\"`"+`
	}

	// Unmarshal the XML into our strongly-typed struct
	if err := xml.Unmarshal(xmlData, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal XML: %%w", err)
	}

	// Parse and execute the template
	tmpl, err := template.New("parser").Parse(%q)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %%w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, response.Body.Response); err != nil {
		return "", fmt.Errorf("failed to execute template: %%w", err)
	}
	return buf.String(), nil
}
`, operationName, templateContent); err != nil {
		return fmt.Errorf("failed to write parser function: %w", err)
	}

	return nil
}

// extractOperationName extracts the operation name from a SOAP endpoint
func extractOperationName(endpoint string) string {
	// Remove any namespace prefixes and URLs
	if idx := strings.LastIndex(endpoint, "/"); idx != -1 {
		endpoint = endpoint[idx+1:]
	}
	if idx := strings.Index(endpoint, ":"); idx != -1 {
		endpoint = endpoint[idx+1:]
	}
	// Remove any query parameters
	if idx := strings.Index(endpoint, "?"); idx != -1 {
		endpoint = endpoint[:idx]
	}
	// Remove any file extensions
	if idx := strings.LastIndex(endpoint, "."); idx != -1 {
		endpoint = endpoint[:idx]
	}
	return endpoint
}
