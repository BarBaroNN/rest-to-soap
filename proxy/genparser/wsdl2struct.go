package genparser

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// WSDLType and WSDLElement are minimal representations for parsing
// Only the fields we need for struct extraction are included

type wsdlDefinitions struct {
	XMLName xml.Name `xml:"definitions"`
	Types   struct {
		Schemas []xsSchema `xml:"schema"`
	} `xml:"types"`
	Messages []wsdlMessage `xml:"message"`
	PortType struct {
		Operations []wsdlOperation `xml:"operation"`
	} `xml:"portType"`
}

type wsdlMessage struct {
	Name  string `xml:"name,attr"`
	Parts []struct {
		Name string `xml:"name,attr"`
		Type string `xml:"type,attr"`
		Elem string `xml:"element,attr"`
	} `xml:"part"`
}

type wsdlOperation struct {
	Name  string `xml:"name,attr"`
	Input struct {
		Message string `xml:"message,attr"`
	} `xml:"input"`
	Output struct {
		Message string `xml:"message,attr"`
	} `xml:"output"`
}

type xsSchema struct {
	XMLName      xml.Name        `xml:"schema"`
	TargetNS     string          `xml:"targetNamespace,attr"`
	ComplexTypes []xsComplexType `xml:"complexType"`
	SimpleTypes  []xsSimpleType  `xml:"simpleType"`
	Elements     []xsElement     `xml:"element"`
	Imports      []xsImport      `xml:"import"`
}

type xsImport struct {
	SchemaLocation string `xml:"schemaLocation,attr"`
	Namespace      string `xml:"namespace,attr"`
}

type xsComplexType struct {
	Name          string           `xml:"name,attr"`
	Sequence      *xsSequence      `xml:"sequence"`
	Attributes    []xsAttribute    `xml:"attribute"`
	SimpleContent *xsSimpleContent `xml:"simpleContent"`
}

type xsSimpleContent struct {
	Extension xsExtension `xml:"extension"`
}

type xsExtension struct {
	Base       string        `xml:"base,attr"`
	Attributes []xsAttribute `xml:"attribute"`
}

type xsSimpleType struct {
	Name        string `xml:"name,attr"`
	Restriction struct {
		Base  string `xml:"base,attr"`
		Enums []struct {
			Value string `xml:"value,attr"`
		} `xml:"enumeration"`
	} `xml:"restriction"`
}

type xsSequence struct {
	Elements []xsElement `xml:"element"`
}

type xsElement struct {
	Name        string         `xml:"name,attr"`
	Type        string         `xml:"type,attr"`
	MinOccurs   string         `xml:"minOccurs,attr"`
	MaxOccurs   string         `xml:"maxOccurs,attr"`
	Nillable    string         `xml:"nillable,attr"`
	Ref         string         `xml:"ref,attr"`
	ComplexType *xsComplexType `xml:"complexType"`
}

type xsAttribute struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Use  string `xml:"use,attr"`
}

// ExtractStructsFromWSDL parses the WSDL and returns Go struct definitions for the specified endpoint
func ExtractStructsFromWSDL(wsdlPath, endpointName string) (string, string, error) {
	fmt.Printf("Reading WSDL file: %s\n", wsdlPath)

	// Read WSDL content
	var data []byte
	var err error

	if strings.HasPrefix(wsdlPath, "http://") || strings.HasPrefix(wsdlPath, "https://") {
		// Handle HTTP URLs
		client := &http.Client{
			Timeout: 30 * time.Second,
		}
		resp, err := client.Get(wsdlPath)
		if err != nil {
			return "", "", fmt.Errorf("failed to fetch WSDL from URL: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", "", fmt.Errorf("failed to fetch WSDL: HTTP %d", resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return "", "", fmt.Errorf("failed to read WSDL response: %w", err)
		}
	} else {
		// Handle local files
		data, err = os.ReadFile(wsdlPath)
		if err != nil {
			return "", "", fmt.Errorf("failed to read WSDL file: %w", err)
		}
	}

	fmt.Printf("Successfully read WSDL file (%d bytes)\n", len(data))

	var wsdl wsdlDefinitions
	if err := xml.Unmarshal(data, &wsdl); err != nil {
		return "", "", fmt.Errorf("failed to parse WSDL: %w", err)
	}
	fmt.Printf("Successfully parsed WSDL\n")

	// Find the operation
	fmt.Printf("Looking for endpoint: %s\n", endpointName)
	var operation wsdlOperation
	for _, op := range wsdl.PortType.Operations {
		fmt.Printf("Found operation: %s\n", op.Name)
		if op.Name == endpointName {
			operation = op
			break
		}
	}
	if operation.Name == "" {
		return "", "", fmt.Errorf("endpoint %s not found", endpointName)
	}
	fmt.Printf("Found matching operation: %s\n", operation.Name)

	// Find the output message
	outputMsgName := operation.Output.Message
	if idx := strings.Index(outputMsgName, ":"); idx != -1 {
		outputMsgName = outputMsgName[idx+1:]
	}
	fmt.Printf("Looking for output message: %s\n", outputMsgName)
	var outputMessage wsdlMessage
	for _, msg := range wsdl.Messages {
		fmt.Printf("Found message: %s\n", msg.Name)
		if msg.Name == outputMsgName {
			outputMessage = msg
			break
		}
	}
	if outputMessage.Name == "" {
		return "", "", fmt.Errorf("output message not found for endpoint %s", endpointName)
	}
	fmt.Printf("Found matching message: %s\n", outputMessage.Name)

	// Get the response element type
	var responseType string
	for _, part := range outputMessage.Parts {
		fmt.Printf("Message part - Name: %s, Type: %s, Elem: %s\n", part.Name, part.Type, part.Elem)
		if part.Elem != "" {
			responseType = part.Elem
			break
		}
	}
	if responseType == "" {
		return "", "", fmt.Errorf("response type not found for endpoint %s", endpointName)
	}
	fmt.Printf("Found response type: %s\n", responseType)

	// Build type and element maps from all schemas
	typeMap := make(map[string]xsComplexType)
	elementMap := make(map[string]xsElement)
	for _, schema := range wsdl.Types.Schemas {
		fmt.Printf("Processing schema with target namespace: %s\n", schema.TargetNS)
		for _, t := range schema.ComplexTypes {
			fmt.Printf("Found complex type: %s\n", t.Name)
			typeMap[t.Name] = t
		}
		for _, e := range schema.Elements {
			fmt.Printf("Found element: %s (type: %s, ref: %s, minOccurs: %s, maxOccurs: %s)\n",
				e.Name, e.Type, e.Ref, e.MinOccurs, e.MaxOccurs)

			// If element has an inline complex type, add it to typeMap
			if e.ComplexType != nil {
				fmt.Printf("Element %s has inline complex type\n", e.Name)
				e.ComplexType.Name = e.Name
				typeMap[e.Name] = *e.ComplexType
			}

			elementMap[e.Name] = e
		}
	}

	// Print all elements and their types for debugging
	fmt.Printf("\nAll elements in elementMap:\n")
	for name, elem := range elementMap {
		fmt.Printf("Element: %s\n", name)
		fmt.Printf("  Type: %s\n", elem.Type)
		fmt.Printf("  Ref: %s\n", elem.Ref)
		fmt.Printf("  MinOccurs: %s\n", elem.MinOccurs)
		fmt.Printf("  MaxOccurs: %s\n", elem.MaxOccurs)
		if elem.ComplexType != nil {
			fmt.Printf("  Has inline complex type\n")
		}
	}

	// Print all complex types for debugging
	fmt.Printf("\nAll complex types in typeMap:\n")
	for name, t := range typeMap {
		fmt.Printf("Complex Type: %s\n", name)
		if t.Sequence != nil {
			fmt.Printf("  Has sequence with %d elements\n", len(t.Sequence.Elements))
			for _, e := range t.Sequence.Elements {
				fmt.Printf("    Element: %s (type: %s, minOccurs: %s, maxOccurs: %s)\n",
					e.Name, e.Type, e.MinOccurs, e.MaxOccurs)
			}
		}
		if len(t.Attributes) > 0 {
			fmt.Printf("  Has %d attributes\n", len(t.Attributes))
			for _, attr := range t.Attributes {
				fmt.Printf("    Attribute: %s (type: %s)\n", attr.Name, attr.Type)
			}
		}
		if t.SimpleContent != nil {
			fmt.Printf("  Has simple content with base: %s\n", t.SimpleContent.Extension.Base)
		}
	}

	// Generate structs recursively
	structs := make(map[string]string)
	visited := make(map[string]bool)
	fmt.Printf("Starting recursive struct generation for type: %s\n", responseType)
	if err := buildStructsRecursive(typeMap, elementMap, responseType, structs, visited, ""); err != nil {
		return "", responseType, fmt.Errorf("failed to build structs: %w", err)
	}

	// Also build structs for all complex types in the typeMap
	fmt.Printf("\nBuilding structs for all complex types:\n")
	for name := range typeMap {
		if !visited[name] {
			fmt.Printf("Building struct for complex type: %s\n", name)
			if err := buildStructsRecursive(typeMap, elementMap, name, structs, visited, ""); err != nil {
				return "", responseType, fmt.Errorf("failed to build struct for complex type %s: %w", name, err)
			}
		}
	}

	// Combine all structs
	var out strings.Builder
	fmt.Printf("Generated %d structs:\n", len(structs))
	for name, s := range structs {
		fmt.Printf("Struct: %s\n%s\n", name, s)
		out.WriteString(s + "\n\n")
	}

	return out.String(), GoTypeName(responseType), nil
}

// buildStructsRecursive generates Go structs for the given type/element name recursively
func buildStructsRecursive(typeMap map[string]xsComplexType, elementMap map[string]xsElement, typeName string, structs map[string]string, visited map[string]bool, _ string) error {
	// Remove namespace
	baseTypeName := typeName
	if idx := strings.Index(baseTypeName, ":"); idx != -1 {
		baseTypeName = baseTypeName[idx+1:]
	}
	fmt.Printf("Building struct for type: %s (base name: %s)\n", typeName, baseTypeName)

	if baseTypeName == "" {
		fmt.Printf("Empty type name, skipping\n")
		return nil
	}

	// Check if it's a built-in type first
	if isBuiltInType(typeName) {
		fmt.Printf("Type %s is a built-in XSD type, skipping struct generation\n", typeName)
		return nil
	}

	// Check if it's a complex type first
	if t, ok := typeMap[baseTypeName]; ok {
		// Only mark as visited if we're actually going to process it
		if !visited[baseTypeName] {
			visited[baseTypeName] = true
			fmt.Printf("Found complex type definition for %s\n", baseTypeName)

			var sb strings.Builder
			structName := GoTypeName(baseTypeName)
			sb.WriteString("type " + structName + " struct {\n")

			// Handle sequence elements
			if t.Sequence != nil {
				fmt.Printf("Processing sequence with %d elements\n", len(t.Sequence.Elements))
				for _, e := range t.Sequence.Elements {
					// Get the actual type, handling both direct types and references
					fieldType := e.Type
					if fieldType == "" && e.Ref != "" {
						fieldType = e.Ref
					}

					// If the element has a complex type definition, use that
					if e.ComplexType != nil {
						// Add the complex type to the type map with a unique name
						complexTypeName := baseTypeName + "_" + e.Name
						typeMap[complexTypeName] = *e.ComplexType
						fieldType = complexTypeName
					}

					// Process the element's type first if it's a complex type and not a built-in type
					if fieldType != "" && !isBuiltInType(fieldType) {
						if err := buildStructsRecursive(typeMap, elementMap, fieldType, structs, visited, e.Name); err != nil {
							return err
						}
					}

					// Convert the type to a Go type
					goType := GoTypeName(fieldType)
					if e.MaxOccurs != "" && e.MaxOccurs != "1" {
						goType = "[]" + goType
					}

					fmt.Printf("Adding sequence element %s of type %s (Go type: %s)\n", e.Name, fieldType, goType)
					sb.WriteString("\t" + goFieldName(e.Name) + " " + goType + " `xml:\"" + e.Name + "\"`\n")
				}
			}

			// Handle attributes
			if len(t.Attributes) > 0 {
				fmt.Printf("Processing %d attributes\n", len(t.Attributes))
				for _, attr := range t.Attributes {
					fieldType := GoTypeName(attr.Type)
					fmt.Printf("Adding attribute %s of type %s\n", attr.Name, fieldType)
					sb.WriteString("\t" + goFieldName(attr.Name) + " " + fieldType + " `xml:\"" + attr.Name + ",attr\"`\n")
				}
			}

			// Handle simple content
			if t.SimpleContent != nil {
				fmt.Printf("Processing simple content with base type %s\n", t.SimpleContent.Extension.Base)
				baseType := GoTypeName(t.SimpleContent.Extension.Base)
				sb.WriteString("\tValue " + baseType + " `xml:\",chardata\"`\n")
				for _, attr := range t.SimpleContent.Extension.Attributes {
					fieldType := GoTypeName(attr.Type)
					fmt.Printf("Adding simple content attribute %s of type %s\n", attr.Name, fieldType)
					sb.WriteString("\t" + goFieldName(attr.Name) + " " + fieldType + " `xml:\"" + attr.Name + ",attr\"`\n")
				}
			}

			sb.WriteString("}")
			structs[structName] = sb.String()
		}
		return nil
	}

	// If not a complex type, check if it's an element
	if elem, ok := elementMap[baseTypeName]; ok {
		fmt.Printf("Found element definition for %s\n", baseTypeName)
		// If the element has a type, use that
		if elem.Type != "" {
			fmt.Printf("Element has type: %s\n", elem.Type)
			// Process the element's type first if it's not a built-in type
			if !isBuiltInType(elem.Type) {
				if err := buildStructsRecursive(typeMap, elementMap, elem.Type, structs, visited, elem.Name); err != nil {
					return err
				}
			}

			// Create a struct for the element
			var sb strings.Builder
			structName := GoTypeName(baseTypeName)
			sb.WriteString("type " + structName + " struct {\n")

			// Add a field for the element's type
			fieldType := GoTypeName(elem.Type)
			if elem.MaxOccurs != "" && elem.MaxOccurs != "1" {
				fieldType = "[]" + fieldType
			}
			fmt.Printf("Adding field %s of type %s\n", elem.Name, fieldType)
			sb.WriteString("\t" + goFieldName(elem.Name) + " " + fieldType + " `xml:\"" + elem.Name + "\"`\n")

			sb.WriteString("}")
			structs[structName] = sb.String()
			return nil
		}

		// If the element references another element
		if elem.Ref != "" {
			fmt.Printf("Element references another element: %s\n", elem.Ref)
			refName := elem.Ref
			if idx := strings.Index(refName, ":"); idx != -1 {
				refName = refName[idx+1:]
			}
			return buildStructsRecursive(typeMap, elementMap, refName, structs, visited, elem.Name)
		}

		// If the element has no type or ref, look for its complex type definition
		fmt.Printf("Element has no type or ref, looking for complex type definition\n")
		var sb strings.Builder
		structName := GoTypeName(baseTypeName)
		sb.WriteString("type " + structName + " struct {\n")

		// Look for the element's complex type in the schema
		foundComplexType := false
		for _, t := range typeMap {
			if t.Name == baseTypeName {
				foundComplexType = true
				fmt.Printf("Found matching complex type for %s\n", baseTypeName)
				// Handle sequence elements
				if t.Sequence != nil {
					fmt.Printf("Processing sequence with %d elements\n", len(t.Sequence.Elements))
					for _, e := range t.Sequence.Elements {
						// Get the actual type, handling both direct types and references
						fieldType := e.Type
						if fieldType == "" && e.Ref != "" {
							fieldType = e.Ref
						}

						// If the element has a complex type definition, use that
						if e.ComplexType != nil {
							// Add the complex type to the type map with a unique name
							complexTypeName := baseTypeName + "_" + e.Name
							typeMap[complexTypeName] = *e.ComplexType
							fieldType = complexTypeName
						}

						// Process the element's type first if it's a complex type and not a built-in type
						if fieldType != "" && !isBuiltInType(fieldType) {
							if err := buildStructsRecursive(typeMap, elementMap, fieldType, structs, visited, e.Name); err != nil {
								return err
							}
						}

						// Convert the type to a Go type
						goType := GoTypeName(fieldType)
						if e.MaxOccurs != "" && e.MaxOccurs != "1" {
							goType = "[]" + goType
						}

						fmt.Printf("Adding sequence element %s of type %s (Go type: %s)\n", e.Name, fieldType, goType)
						sb.WriteString("\t" + goFieldName(e.Name) + " " + goType + " `xml:\"" + e.Name + "\"`\n")
					}
				}

				// Handle attributes
				if len(t.Attributes) > 0 {
					fmt.Printf("Processing %d attributes\n", len(t.Attributes))
					for _, attr := range t.Attributes {
						fieldType := GoTypeName(attr.Type)
						fmt.Printf("Adding attribute %s of type %s\n", attr.Name, fieldType)
						sb.WriteString("\t" + goFieldName(attr.Name) + " " + fieldType + " `xml:\"" + attr.Name + ",attr\"`\n")
					}
				}

				// Handle simple content
				if t.SimpleContent != nil {
					fmt.Printf("Processing simple content with base type %s\n", t.SimpleContent.Extension.Base)
					baseType := GoTypeName(t.SimpleContent.Extension.Base)
					sb.WriteString("\tValue " + baseType + " `xml:\",chardata\"`\n")
					for _, attr := range t.SimpleContent.Extension.Attributes {
						fieldType := GoTypeName(attr.Type)
						fmt.Printf("Adding simple content attribute %s of type %s\n", attr.Name, fieldType)
						sb.WriteString("\t" + goFieldName(attr.Name) + " " + fieldType + " `xml:\"" + attr.Name + ",attr\"`\n")
					}
				}
			}
		}

		if !foundComplexType {
			// If no complex type found, try to find a matching element with the same name
			if matchingElem, ok := elementMap[baseTypeName]; ok {
				fmt.Printf("Found matching element for %s\n", baseTypeName)
				if matchingElem.Type != "" {
					// Process the element's type first if it's not a built-in type
					if !isBuiltInType(matchingElem.Type) {
						if err := buildStructsRecursive(typeMap, elementMap, matchingElem.Type, structs, visited, matchingElem.Name); err != nil {
							return err
						}
					}

					fieldType := GoTypeName(matchingElem.Type)
					if matchingElem.MaxOccurs != "" && matchingElem.MaxOccurs != "1" {
						fieldType = "[]" + fieldType
					}
					fmt.Printf("Adding field %s of type %s\n", matchingElem.Name, fieldType)
					sb.WriteString("\t" + goFieldName(matchingElem.Name) + " " + fieldType + " `xml:\"" + matchingElem.Name + "\"`\n")
				}
			} else {
				fmt.Printf("Warning: No complex type or matching element found for %s\n", baseTypeName)
			}
		}

		sb.WriteString("}")
		structs[structName] = sb.String()
		return nil
	}

	return fmt.Errorf("type %s not found in type map or element map", baseTypeName)
}

// GoTypeName converts an XSD type name to a Go type name
func GoTypeName(xsdType string) string {
	if idx := strings.Index(xsdType, ":"); idx != -1 {
		xsdType = xsdType[idx+1:]
	}
	switch xsdType {
	case "string", "normalizedString", "token":
		return "string"
	case "int", "integer", "long", "short", "byte":
		return "int"
	case "decimal", "float", "double":
		return "float64"
	case "boolean":
		return "bool"
	case "date", "dateTime", "time":
		return "time.Time"
	case "base64Binary":
		return "[]byte"
	case "anyURI":
		return "string"
	case "QName":
		return "string"
	}
	return xsdType
}

// isBuiltInType checks if the type is a built-in XSD type
func isBuiltInType(typeName string) bool {
	if idx := strings.Index(typeName, ":"); idx != -1 {
		typeName = typeName[idx+1:]
	}
	switch typeName {
	case "string", "normalizedString", "token",
		"int", "integer", "long", "short", "byte",
		"decimal", "float", "double",
		"boolean",
		"date", "dateTime", "time",
		"base64Binary",
		"anyURI",
		"QName":
		return true
	}
	return false
}

func goFieldName(xmlName string) string {
	if len(xmlName) == 0 {
		return ""
	}
	return strings.ToUpper(xmlName[:1]) + xmlName[1:]
}
