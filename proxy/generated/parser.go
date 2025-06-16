package generated

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

type tLanguage struct {
	SISOCode string `xml:"sISOCode"`
	SName    string `xml:"sName"`
}

type ArrayOftLanguage struct {
	TLanguage []tLanguage `xml:"tLanguage"`
}

type LanguageName struct {
	SISOCode string `xml:"sISOCode"`
}

type ListOfCurrenciesByCodeResponse struct {
	ListOfCurrenciesByCodeResult ArrayOftCurrency `xml:"ListOfCurrenciesByCodeResult"`
}

type ArrayOftCountryCodeAndName struct {
	TCountryCodeAndName []tCountryCodeAndName `xml:"tCountryCodeAndName"`
}

type CountryIntPhoneCodeResponse struct {
	CountryIntPhoneCodeResult string `xml:"CountryIntPhoneCodeResult"`
}

type ListOfCurrenciesByCode struct {
}

type CapitalCityResponse struct {
	CapitalCityResult string `xml:"CapitalCityResult"`
}

type CountriesUsingCurrency struct {
	SISOCurrencyCode string `xml:"sISOCurrencyCode"`
}

type tCountryCodeAndName struct {
	SISOCode string `xml:"sISOCode"`
	SName    string `xml:"sName"`
}

type tCurrency struct {
	SISOCode string `xml:"sISOCode"`
	SName    string `xml:"sName"`
}

type ListOfLanguagesByCode struct {
}

type LanguageISOCodeResponse struct {
	LanguageISOCodeResult string `xml:"LanguageISOCodeResult"`
}

type FullCountryInfoResponse struct {
	FullCountryInfoResult tCountryInfo `xml:"FullCountryInfoResult"`
}

type FullCountryInfoAllCountries struct {
}

type ListOfCountryNamesGroupedByContinent struct {
}

type CountryISOCodeResponse struct {
	CountryISOCodeResult string `xml:"CountryISOCodeResult"`
}

type CountryCurrencyResponse struct {
	CountryCurrencyResult tCurrency `xml:"CountryCurrencyResult"`
}

type FullCountryInfo struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ArrayOftContinent struct {
	TContinent []tContinent `xml:"tContinent"`
}

type ArrayOftCountryCodeAndNameGroupedByContinent struct {
	TCountryCodeAndNameGroupedByContinent []tCountryCodeAndNameGroupedByContinent `xml:"tCountryCodeAndNameGroupedByContinent"`
}

type ListOfContinentsByCodeResponse struct {
	ListOfContinentsByCodeResult ArrayOftContinent `xml:"ListOfContinentsByCodeResult"`
}

type FullCountryInfoAllCountriesResponse struct {
	FullCountryInfoAllCountriesResult ArrayOftCountryInfo `xml:"FullCountryInfoAllCountriesResult"`
}

type ListOfLanguagesByName struct {
}

type ArrayOftCurrency struct {
	TCurrency []tCurrency `xml:"tCurrency"`
}

type ListOfCountryNamesByCodeResponse struct {
	ListOfCountryNamesByCodeResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByCodeResult"`
}

type ListOfCountryNamesByNameResponse struct {
	ListOfCountryNamesByNameResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByNameResult"`
}

type LanguageISOCode struct {
	SLanguageName string `xml:"sLanguageName"`
}

type ListOfCountryNamesGroupedByContinentResponse struct {
	ListOfCountryNamesGroupedByContinentResult ArrayOftCountryCodeAndNameGroupedByContinent `xml:"ListOfCountryNamesGroupedByContinentResult"`
}

type LanguageNameResponse struct {
	LanguageNameResult string `xml:"LanguageNameResult"`
}

type ListOfContinentsByCode struct {
}

type ListOfCurrenciesByName struct {
}

type ListOfCountryNamesByName struct {
}

type ListOfCurrenciesByNameResponse struct {
	ListOfCurrenciesByNameResult ArrayOftCurrency `xml:"ListOfCurrenciesByNameResult"`
}

type CountryCurrency struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type CapitalCity struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type CurrencyName struct {
	SCurrencyISOCode string `xml:"sCurrencyISOCode"`
}

type ListOfLanguagesByNameResponse struct {
	ListOfLanguagesByNameResult ArrayOftLanguage `xml:"ListOfLanguagesByNameResult"`
}

type CountriesUsingCurrencyResponse struct {
	CountriesUsingCurrencyResult ArrayOftCountryCodeAndName `xml:"CountriesUsingCurrencyResult"`
}

type ListOfContinentsByNameResponse struct {
	ListOfContinentsByNameResult ArrayOftContinent `xml:"ListOfContinentsByNameResult"`
}

type tCountryInfo struct {
	SISOCode         string           `xml:"sISOCode"`
	SName            string           `xml:"sName"`
	SCapitalCity     string           `xml:"sCapitalCity"`
	SPhoneCode       string           `xml:"sPhoneCode"`
	SContinentCode   string           `xml:"sContinentCode"`
	SCurrencyISOCode string           `xml:"sCurrencyISOCode"`
	SCountryFlag     string           `xml:"sCountryFlag"`
	Languages        ArrayOftLanguage `xml:"Languages"`
}

type ArrayOftCountryInfo struct {
	TCountryInfo []tCountryInfo `xml:"tCountryInfo"`
}

type tCountryCodeAndNameGroupedByContinent struct {
	Continent           tContinent                 `xml:"Continent"`
	CountryCodeAndNames ArrayOftCountryCodeAndName `xml:"CountryCodeAndNames"`
}

type CountryISOCode struct {
	SCountryName string `xml:"sCountryName"`
}

type CountryFlag struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type CurrencyNameResponse struct {
	CurrencyNameResult string `xml:"CurrencyNameResult"`
}

type CountryIntPhoneCode struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type CountryNameResponse struct {
	CountryNameResult string `xml:"CountryNameResult"`
}

type ListOfLanguagesByCodeResponse struct {
	ListOfLanguagesByCodeResult ArrayOftLanguage `xml:"ListOfLanguagesByCodeResult"`
}

type CountryName struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ListOfContinentsByName struct {
}

type ListOfCountryNamesByCode struct {
}

type tContinent struct {
	SCode string `xml:"sCode"`
	SName string `xml:"sName"`
}

type CountryFlagResponse struct {
	CountryFlagResult string `xml:"CountryFlagResult"`
}

// FullCountryInfoAllCountriesParser parses the SOAP response for the FullCountryInfoAllCountries operation
func Parse(xmlData []byte) (string, error) {
	// Define the SOAP envelope structure with the proper response type
	var response struct {
		XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Body    struct {
			Response FullCountryInfoAllCountriesResponse `xml:"FullCountryInfoAllCountriesResponse"`
		} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	}

	// Unmarshal the XML into our strongly-typed struct
	if err := xml.Unmarshal(xmlData, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("templates/response.tmpl")
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, response.Body.Response); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}
