package genparser

import (
	"fmt"
	"os"
	"path/filepath"
	"rest-to-soap/proxy/config"
)

type RegistryGenerator struct {
	outputDir string
}

func NewRegistryGenerator() *RegistryGenerator {
	return &RegistryGenerator{
		outputDir: "generated",
	}
}

func (g *RegistryGenerator) GenerateRegistry(cfg *config.Config) error {
	outputPath := filepath.Join(g.outputDir, "route_handler_registry.go")
	routeHandlers := generateRouteHandlers(cfg)
	generatedCode := fmt.Sprintf(`
	package generated

	import (
		"rest-to-soap/proxy/config"
		"text/template"
	)

	type GeneratedRouteHandler struct {
		RouteConfig      config.RouteConfig
		Parser           func([]byte) (string, error)
		RequestTemplate  template.Template
		ResponseTemplate template.Template
	}

	type RouteRegistry map[string]GeneratedRouteHandler

	var RouteHandlerRegistry = RouteRegistry{
		%s
	}`, routeHandlers)

	return os.WriteFile(outputPath, []byte(generatedCode), 0644)
}

func generateRouteHandlers(cfg *config.Config) string {
	generatedCode := ""

	for _, route := range cfg.Routes {
		generatedCode += fmt.Sprintf(`
			"%s": {
				RouteConfig: %v,
				Parser:      %sParse,
				RequestTemplate: template.Template{},
				ResponseTemplate: template.Template{},
			},
		`, route.Path, fmt.Sprintf("%#v", route), route.SoapAction)
	}

	return generatedCode
}
