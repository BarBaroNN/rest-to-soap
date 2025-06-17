
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
		
			"/api/soap/countries": {
				RouteConfig: config.RouteConfig{Path:"/api/soap/countries", Method:"POST", SoapEndpoint:"http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso", SoapAction:"CountryFlag", RequestTemplate:"templates/request.tmpl", ResponseTemplate:"templates/response.tmpl", Headers:map[string]string{"Content-Type":"text/xml;charset=UTF-8", "SOAPAction":"CountryFlag"}, WSDLURL:"http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso?WSDL", Timeout:30000000000},
				Parser:      CountryFlagParse,
				RequestTemplate: template.Template{},
				ResponseTemplate: template.Template{},
			},
		
			"/api/soap/degrees/celsius-to-fahrenheit": {
				RouteConfig: config.RouteConfig{Path:"/api/soap/degrees/celsius-to-fahrenheit", Method:"POST", SoapEndpoint:"https://www.w3schools.com/xml/tempconvert.asmx", SoapAction:"CelsiusToFahrenheit", RequestTemplate:"templates/celsius-to-farenheit-request.tmpl", ResponseTemplate:"templates/celsius-to-farenheit-response.tmpl", Headers:map[string]string{"Content-Type":"text/xml;charset=UTF-8"}, WSDLURL:"https://www.w3schools.com/xml/tempconvert.asmx?WSDL", Timeout:30000000000},
				Parser:      CelsiusToFahrenheitParse,
				RequestTemplate: template.Template{},
				ResponseTemplate: template.Template{},
			},
		
	}