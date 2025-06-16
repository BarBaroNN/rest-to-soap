package ..

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

// Generated structs from WSDL
type CapitalCityResponse struct {
	CapitalCityResult  `xml:"CapitalCityResult"`
}

type ListOfLanguagesByCodeResponse struct {
	ListOfLanguagesByCodeResult  `xml:"ListOfLanguagesByCodeResult"`
}

type ListOfLanguagesByNameResponse struct {
	ListOfLanguagesByNameResult  `xml:"ListOfLanguagesByNameResult"`
}

type LanguageNameResponse struct {
	LanguageNameResult  `xml:"LanguageNameResult"`
}

type ArrayOftLanguage struct {
	TLanguage [] `xml:"tLanguage"`
}

type tCountryCodeAndNameGroupedByContinent struct {
	Continent  `xml:"Continent"`
	CountryCodeAndNames  `xml:"CountryCodeAndNames"`
}

type ListOfCurrenciesByCode struct {
}

type FullCountryInfo struct {
	SCountryISOCode  `xml:"sCountryISOCode"`
}

type ListOfContinentsByNameResponse struct {
	ListOfContinentsByNameResult  `xml:"ListOfContinentsByNameResult"`
}

type CountryISOCodeResponse struct {
	CountryISOCodeResult  `xml:"CountryISOCodeResult"`
}

type ArrayOftContinent struct {
	TContinent [] `xml:"tContinent"`
}

type CurrencyNameResponse struct {
	CurrencyNameResult  `xml:"CurrencyNameResult"`
}

type FullCountryInfoResponse struct {
	FullCountryInfoResult  `xml:"FullCountryInfoResult"`
}

type CountryISOCode struct {
	SCountryName  `xml:"sCountryName"`
}

type CountriesUsingCurrencyResponse struct {
	CountriesUsingCurrencyResult  `xml:"CountriesUsingCurrencyResult"`
}

type tLanguage struct {
	SISOCode  `xml:"sISOCode"`
	SName  `xml:"sName"`
}

type CountryFlag struct {
	SCountryISOCode  `xml:"sCountryISOCode"`
}

type CountriesUsingCurrency struct {
	SISOCurrencyCode  `xml:"sISOCurrencyCode"`
}

type LanguageName struct {
	SISOCode  `xml:"sISOCode"`
}

type ListOfCountryNamesByCode struct {
}

type ListOfLanguagesByCode struct {
}

type FullCountryInfoAllCountries struct {
}

type ListOfCurrenciesByCodeResponse struct {
	ListOfCurrenciesByCodeResult  `xml:"ListOfCurrenciesByCodeResult"`
}

type CapitalCity struct {
	SCountryISOCode  `xml:"sCountryISOCode"`
}

type tCountryInfo struct {
	SISOCode  `xml:"sISOCode"`
	SName  `xml:"sName"`
	SCapitalCity  `xml:"sCapitalCity"`
	SPhoneCode  `xml:"sPhoneCode"`
	SContinentCode  `xml:"sContinentCode"`
	SCurrencyISOCode  `xml:"sCurrencyISOCode"`
	SCountryFlag  `xml:"sCountryFlag"`
	Languages  `xml:"Languages"`
}

type ListOfContinentsByName struct {
}

type FullCountryInfoAllCountriesResponse struct {
	FullCountryInfoAllCountriesResult  `xml:"FullCountryInfoAllCountriesResult"`
}

type ListOfCurrenciesByNameResponse struct {
	ListOfCurrenciesByNameResult  `xml:"ListOfCurrenciesByNameResult"`
}

type CountryFlagResponse struct {
	CountryFlagResult  `xml:"CountryFlagResult"`
}

type ListOfLanguagesByName struct {
}

type tContinent struct {
	SCode  `xml:"sCode"`
	SName  `xml:"sName"`
}

type CountryIntPhoneCodeResponse struct {
	CountryIntPhoneCodeResult  `xml:"CountryIntPhoneCodeResult"`
}

type ArrayOftCountryCodeAndName struct {
	TCountryCodeAndName [] `xml:"tCountryCodeAndName"`
}

type CountryNameResponse struct {
	CountryNameResult  `xml:"CountryNameResult"`
}

type CountryCurrencyResponse struct {
	CountryCurrencyResult  `xml:"CountryCurrencyResult"`
}

type ArrayOftCurrency struct {
	TCurrency [] `xml:"tCurrency"`
}

type ListOfCountryNamesByName struct {
}

type CountryCurrency struct {
	SCountryISOCode  `xml:"sCountryISOCode"`
}

type tCountryCodeAndName struct {
	SISOCode  `xml:"sISOCode"`
	SName  `xml:"sName"`
}

type ListOfContinentsByCode struct {
}

type ArrayOftCountryInfo struct {
	TCountryInfo [] `xml:"tCountryInfo"`
}

type ListOfCountryNamesGroupedByContinentResponse struct {
	ListOfCountryNamesGroupedByContinentResult  `xml:"ListOfCountryNamesGroupedByContinentResult"`
}

type CountryName struct {
	SCountryISOCode  `xml:"sCountryISOCode"`
}

type LanguageISOCodeResponse struct {
	LanguageISOCodeResult  `xml:"LanguageISOCodeResult"`
}

type LanguageISOCode struct {
	SLanguageName  `xml:"sLanguageName"`
}

type ArrayOftCountryCodeAndNameGroupedByContinent struct {
	TCountryCodeAndNameGroupedByContinent [] `xml:"tCountryCodeAndNameGroupedByContinent"`
}

type ListOfCountryNamesByNameResponse struct {
	ListOfCountryNamesByNameResult  `xml:"ListOfCountryNamesByNameResult"`
}

type tCurrency struct {
	SISOCode  `xml:"sISOCode"`
	SName  `xml:"sName"`
}

type ListOfCountryNamesByCodeResponse struct {
	ListOfCountryNamesByCodeResult  `xml:"ListOfCountryNamesByCodeResult"`
}

type ListOfContinentsByCodeResponse struct {
	ListOfContinentsByCodeResult  `xml:"ListOfContinentsByCodeResult"`
}

type CountryIntPhoneCode struct {
	SCountryISOCode  `xml:"sCountryISOCode"`
}

type CurrencyName struct {
	SCurrencyISOCode  `xml:"sCurrencyISOCode"`
}

type ListOfCurrenciesByName struct {
}

type ListOfCountryNamesGroupedByContinent struct {
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