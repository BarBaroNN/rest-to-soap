package parser

import (
	"fmt"
	"os"
	"strings"

	"rest-to-soap/proxy/generated"
)

// ParserRegistry maintains a registry of SOAP response parsers
type ParserRegistry struct {
	parsers map[string]func([]byte) (string, error)
}

// NewParserRegistry creates a new parser registry
func NewParserRegistry() *ParserRegistry {
	return &ParserRegistry{
		parsers: make(map[string]func([]byte) (string, error)),
	}
}

// RegisterParser registers a parser function for a specific SOAP action
func (r *ParserRegistry) RegisterParser(soapAction string, parser func([]byte) (string, error)) {
	r.parsers[soapAction] = parser
}

// LoadGeneratedParsers loads all generated parsers from the generated directory
func (r *ParserRegistry) LoadGeneratedParsers() error {
	// Read the generated directory
	files, err := os.ReadDir("generated")
	if err != nil {
		return fmt.Errorf("failed to read generated directory: %w", err)
	}

	// For each .go file, extract the SOAP action and register its parser
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".go") {
			continue
		}

		soapAction := strings.TrimSuffix(file.Name(), ".go")

		// Get the parser function from the generated package
		parserName := soapAction + "Parser"
		parser, err := getParserFunction(parserName)
		if err != nil {
			return fmt.Errorf("failed to get parser for %s: %w", soapAction, err)
		}

		// Register the parser
		r.RegisterParser(soapAction, parser)
	}
	return nil
}

// Parse parses a SOAP response using the appropriate parser
func (r *ParserRegistry) Parse(soapAction string, xmlData []byte) (string, error) {
	parser, exists := r.parsers[soapAction]
	if !exists {
		return "", fmt.Errorf("no parser found for SOAP action: %s", soapAction)
	}
	return parser(xmlData)
}

// getParserFunction returns the parser function for a given operation
func getParserFunction(parserName string) (func([]byte) (string, error), error) {
	// Use reflection to get the parser function from the generated package
	parser, err := generated.GetParser(parserName)
	if err != nil {
		return nil, fmt.Errorf("failed to get parser %s: %w", parserName, err)
	}
	return parser, nil
}
