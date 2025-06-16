package genparser

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func Generate() {
	wsdlPath := flag.String("wsdl", "", "Path to WSDL file")
	tmplPath := flag.String("tmpl", "", "Path to template file")
	outPath := flag.String("out", "", "Path to output Go file")
	parserName := flag.String("name", "Parser", "Name of the parser function")
	endpointName := flag.String("endpoint", "", "Name of the WSDL endpoint")
	flag.Parse()

	if *wsdlPath == "" || *tmplPath == "" || *outPath == "" || *endpointName == "" {
		fmt.Println("Usage: genparser -wsdl <wsdl> -tmpl <tmpl> -out <out> -endpoint <endpoint> [-name <name>]")
		os.Exit(1)
	}

	// Extract structs from WSDL
	structs, err := ExtractStructsFromWSDL(*wsdlPath, *endpointName)
	if err != nil {
		fmt.Printf("Failed to extract structs from WSDL: %v\n", err)
		os.Exit(1)
	}

	// Load template
	tmplData, err := os.ReadFile(*tmplPath)
	if err != nil {
		fmt.Printf("Failed to read template file: %v\n", err)
		os.Exit(1)
	}

	// Get package name from output path
	pkgName := filepath.Base(filepath.Dir(*outPath))
	if pkgName == "." {
		pkgName = "generated"
	}

	// Generate Go file with proper struct types and SOAP envelope handling
	code := fmt.Sprintf(`package %s

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

// Generated structs from WSDL
%s

// %s parses the SOAP response and renders JSON using the template
func %s(xmlData []byte) (string, error) {
	// Define the SOAP envelope structure with the proper response type
	var response struct {
		XMLName xml.Name `+"`xml:\"http://schemas.xmlsoap.org/soap/envelope/ Envelope\"`"+`
		Body    struct {
			Response %s `+"`xml:\"Response\"`"+`
		} `+"`xml:\"http://schemas.xmlsoap.org/soap/envelope/ Body\"`"+`
	}

	// Unmarshal the XML into our strongly-typed struct
	if err := xml.Unmarshal(xmlData, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal XML: %%w", err)
	}

	// Parse and execute the template with the full response struct
	tmpl, err := template.New("parser").Parse(%q)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %%w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, response.Body.Response); err != nil {
		return "", fmt.Errorf("failed to execute template: %%w", err)
	}
	return buf.String(), nil
}`, pkgName, structs, *parserName, *parserName, GoTypeName(*endpointName+"Response"), string(tmplData))

	if err := os.WriteFile(*outPath, []byte(code), 0644); err != nil {
		fmt.Printf("Failed to write output file: %v\n", err)
		os.Exit(1)
	}
}
