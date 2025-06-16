package generated

import (
	"fmt"
)

var parsers = make(map[string]func([]byte) (string, error))

// RegisterParser registers a parser function
func RegisterParser(name string, parser func([]byte) (string, error)) {
	parsers[name] = parser
}

// GetParser returns a parser function by name
func GetParser(name string) (func([]byte) (string, error), error) {
	parser, exists := parsers[name]
	if !exists {
		return nil, fmt.Errorf("parser %s not found", name)
	}
	return parser, nil
}

// init registers all parser functions
func init() {
	// Register all parser functions
	RegisterParser("FullCountryInfoAllCountriesParser", FullCountryInfoAllCountriesParser)
	// Add more parsers here as they are generated
}

// Parser parses the SOAP response and renders JSON using the template
