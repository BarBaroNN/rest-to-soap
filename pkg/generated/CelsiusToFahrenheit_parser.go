package generated

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

type CelsiusToFahrenheit struct {
	Celsius string `xml:"Celsius"`
}

type CelsiusToFahrenheitResponse struct {
	CelsiusToFahrenheitResult string `xml:"CelsiusToFahrenheitResult"`
}

type FahrenheitToCelsius struct {
	Fahrenheit string `xml:"Fahrenheit"`
}

type FahrenheitToCelsiusResponse struct {
	FahrenheitToCelsiusResult string `xml:"FahrenheitToCelsiusResult"`
}



// CelsiusToFahrenheitParser parses the SOAP response for the CelsiusToFahrenheit operation
func CelsiusToFahrenheitParse(xmlData []byte) (string, error) {
	fmt.Println(string(xmlData))

	// Define the SOAP envelope structure with the proper response type
	var response struct {
		XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Body    struct {
			Response CelsiusToFahrenheitResponse `xml:"CelsiusToFahrenheitResponse"`
		} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	}

	// Unmarshal the XML into our strongly-typed struct
	if err := xml.Unmarshal(xmlData, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("config/templates/celsius-to-farenheit-response.tmpl")
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, response.Body.Response); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	return buf.String(), nil
}
