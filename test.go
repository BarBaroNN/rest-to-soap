package ..

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

// Generated structs from WSDL
type tContinent struct {
	SCode string `xml:"sCode"`
	SName string `xml:"sName"`
}

type CountriesUsingCurrency struct {
	SISOCurrencyCode string `xml:"sISOCurrencyCode"`
}

type CountryISOCode struct {
	SCountryName string `xml:"sCountryName"`
}

type LanguageISOCode struct {
	SLanguageName string `xml:"sLanguageName"`
}

type ListOfCountryNamesByCode struct {
}

type ListOfCountryNamesByCodeResponse struct {
	ListOfCountryNamesByCodeResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByCodeResult"`
}

type ListOfCountryNamesByNameResponse struct {
	ListOfCountryNamesByNameResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByNameResult"`
}

type ArrayOftCurrency struct {
	TCurrency []tCurrency `xml:"tCurrency"`
}

type CountryISOCodeResponse struct {
	CountryISOCodeResult string `xml:"CountryISOCodeResult"`
}

type CountryFlagResponse struct {
	CountryFlagResult string `xml:"CountryFlagResult"`
}

type tCountryCodeAndName struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type ListOfCountryNamesGroupedByContinentResponse struct {
	ListOfCountryNamesGroupedByContinentResult ArrayOftCountryCodeAndNameGroupedByContinent `xml:"ListOfCountryNamesGroupedByContinentResult"`
}

type ListOfCurrenciesByNameResponse struct {
	ListOfCurrenciesByNameResult ArrayOftCurrency `xml:"ListOfCurrenciesByNameResult"`
}

type ListOfCurrenciesByName struct {
}

type FullCountryInfoAllCountriesResponse struct {
	FullCountryInfoAllCountriesResult ArrayOftCountryInfo `xml:"FullCountryInfoAllCountriesResult"`
}

type tCurrency struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type ListOfCountryNamesGroupedByContinent struct {
}

type ArrayOftCountryCodeAndNameGroupedByContinent struct {
	TCountryCodeAndNameGroupedByContinent []tCountryCodeAndNameGroupedByContinent `xml:"tCountryCodeAndNameGroupedByContinent"`
}

type ArrayOftContinent struct {
	TContinent []tContinent `xml:"tContinent"`
}

type CountryNameResponse struct {
	CountryNameResult string `xml:"CountryNameResult"`
}

type ListOfContinentsByCode struct {
}

type CountryCurrency struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ListOfContinentsByCodeResponse struct {
	ListOfContinentsByCodeResult ArrayOftContinent `xml:"ListOfContinentsByCodeResult"`
}

type ListOfCurrenciesByCode struct {
}

type ListOfLanguagesByNameResponse struct {
	ListOfLanguagesByNameResult ArrayOftLanguage `xml:"ListOfLanguagesByNameResult"`
}

type CountryIntPhoneCodeResponse struct {
	CountryIntPhoneCodeResult string `xml:"CountryIntPhoneCodeResult"`
}

type CurrencyNameResponse struct {
	CurrencyNameResult string `xml:"CurrencyNameResult"`
}

type ArrayOftCountryCodeAndName struct {
	TCountryCodeAndName []tCountryCodeAndName `xml:"tCountryCodeAndName"`
}

type LanguageISOCodeResponse struct {
	LanguageISOCodeResult string `xml:"LanguageISOCodeResult"`
}

type CountriesUsingCurrencyResponse struct {
	CountriesUsingCurrencyResult ArrayOftCountryCodeAndName `xml:"CountriesUsingCurrencyResult"`
}

type ListOfLanguagesByName struct {
}

type FullCountryInfo struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type tLanguage struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type ListOfCurrenciesByCodeResponse struct {
	ListOfCurrenciesByCodeResult ArrayOftCurrency `xml:"ListOfCurrenciesByCodeResult"`
}

type CountryFlag struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type CountryCurrencyResponse struct {
	CountryCurrencyResult tCurrency `xml:"CountryCurrencyResult"`
}

type ListOfLanguagesByCode struct {
}

type tCountryCodeAndNameGroupedByContinent struct {
	Continent tContinent `xml:"Continent"`
	CountryCodeAndNames ArrayOftCountryCodeAndName `xml:"CountryCodeAndNames"`
}

type LanguageNameResponse struct {
	LanguageNameResult string `xml:"LanguageNameResult"`
}

type ListOfContinentsByNameResponse struct {
	ListOfContinentsByNameResult ArrayOftContinent `xml:"ListOfContinentsByNameResult"`
}

type CountryIntPhoneCode struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type FullCountryInfoAllCountries struct {
}

type LanguageName struct {
	SISOCode string `xml:"sISOCode"`
}

type CapitalCity struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type FullCountryInfoResponse struct {
	FullCountryInfoResult tCountryInfo `xml:"FullCountryInfoResult"`
}

type ArrayOftLanguage struct {
	TLanguage []tLanguage `xml:"tLanguage"`
}

type ArrayOftCountryInfo struct {
	TCountryInfo []tCountryInfo `xml:"tCountryInfo"`
}

type CountryName struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ListOfLanguagesByCodeResponse struct {
	ListOfLanguagesByCodeResult ArrayOftLanguage `xml:"ListOfLanguagesByCodeResult"`
}

type ListOfContinentsByName struct {
}

type CurrencyName struct {
	SCurrencyISOCode string `xml:"sCurrencyISOCode"`
}

type ListOfCountryNamesByName struct {
}

type CapitalCityResponse struct {
	CapitalCityResult string `xml:"CapitalCityResult"`
}

type tCountryInfo struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
	SCapitalCity string `xml:"sCapitalCity"`
	SPhoneCode string `xml:"sPhoneCode"`
	SContinentCode string `xml:"sContinentCode"`
	SCurrencyISOCode string `xml:"sCurrencyISOCode"`
	SCountryFlag string `xml:"sCountryFlag"`
	Languages ArrayOftLanguage `xml:"Languages"`
}



// Parser parses the SOAP response and renders JSON using the template
func Parser(xmlData []byte) (string, error) {
	// Define the SOAP envelope structure with the proper response type
	var response struct {
		XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Body    struct {
			Response FullCountryInfoAllCountriesResponse `xml:"Response"`
		} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	}

	// Unmarshal the XML into our strongly-typed struct
	if err := xml.Unmarshal(xmlData, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	// Parse and execute the template with the full response struct
	tmpl, err := template.New("parser").Parse("[\r\n{{- $countries := .Envelope.Body.FullCountryInfoAllCountriesResponse.FullCountryInfoAllCountriesResult.tCountryInfo -}}\r\n{{- range $i, $c := $countries }}\r\n  {{- if $i}},\r\n  {{end}}\r\n  {\r\n    \"ISOCode\": \"{{$c.sISOCode}}\",\r\n    \"name\": \"{{$c.sName}}\",\r\n    \"capitalCity\": \"{{$c.sCapitalCity}}\",\r\n    \"phoneCode\": \"{{$c.sPhoneCode}}\",\r\n    \"continentCode\": \"{{$c.sContinentCode}}\",\r\n    \"currencyISOCode\": \"{{$c.sCurrencyISOCode}}\",\r\n    \"countryFlag\": \"{{$c.sCountryFlag}}\",\r\n    \"languages\": [\r\n      {{- $langs := $c.Languages.tLanguage -}}\r\n      {{- range $j, $lang := $langs }}\r\n        {{- if $j}},{{end}}\r\n        {\r\n          \"ISOCode\": \"{{$lang.sISOCode}}\",\r\n          \"name\": \"{{$lang.sName}}\"\r\n        }\r\n      {{- end }}\r\n    ]\r\n  }\r\n{{- end }}\r\n] ")
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, response.Body.Response); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	return buf.String(), nil
}