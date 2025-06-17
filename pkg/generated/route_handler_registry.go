package generated

import (
	"rest-to-soap/core/config"
	"text/template"

	"go.uber.org/zap"
)

type GeneratedRouteHandler struct {
	RouteConfig      config.RouteConfig
	Parser           func([]byte) (string, error)
	RequestTemplate  template.Template
	ResponseTemplate template.Template
}

type RouteRegistry map[string]GeneratedRouteHandler

var RouteHandlerRegistry = RouteRegistry{

	"/api/soap/countries": {
		RouteConfig:      config.RouteConfig{Path: "/api/soap/countries", Method: "POST", SoapEndpoint: "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso", SoapAction: "CountryFlag", RequestTemplate: "config/templates/request.tmpl", ResponseTemplate: "config/templates/response.tmpl", Headers: map[string]string{"Content-Type": "text/xml;charset=UTF-8", "SOAPAction": "CountryFlag"}, WSDLURL: "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso?WSDL", Timeout: 30000000000},
		Parser:           CountryFlagParse,
		RequestTemplate:  template.Template{},
		ResponseTemplate: template.Template{},
	},

	"/api/soap/degrees/celsius-to-fahrenheit": {
		RouteConfig:      config.RouteConfig{Path: "/api/soap/degrees/celsius-to-fahrenheit", Method: "POST", SoapEndpoint: "https://www.w3schools.com/xml/tempconvert.asmx", SoapAction: "CelsiusToFahrenheit", RequestTemplate: "config/templates/celsius-to-farenheit-request.tmpl", ResponseTemplate: "config/templates/celsius-to-farenheit-response.tmpl", Headers: map[string]string{"Content-Type": "text/xml;charset=UTF-8"}, WSDLURL: "https://www.w3schools.com/xml/tempconvert.asmx?WSDL", Timeout: 30000000000},
		Parser:           CelsiusToFahrenheitParse,
		RequestTemplate:  template.Template{},
		ResponseTemplate: template.Template{},
	},
}

// Hydrate the route registry with the templates
func GenerateRouteRegistry(cfg *config.Config, logger *zap.Logger) (RouteRegistry, error) {
	for _, route := range cfg.Routes {
		requestTmpl, err := template.ParseFiles(route.RequestTemplate)
		if err != nil {
			return nil, err
		}

		responseTmpl, err := template.ParseFiles(route.ResponseTemplate)
		if err != nil {
			return nil, err
		}

		RouteHandlerRegistry[route.Path] = GeneratedRouteHandler{
			RouteConfig:      route,
			Parser:           RouteHandlerRegistry[route.Path].Parser,
			RequestTemplate:  *requestTmpl,
			ResponseTemplate: *responseTmpl,
		}
	}

	return RouteHandlerRegistry, nil
}
