package generated

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

type CountriesUsingCurrency struct {
	SISOCurrencyCode string `xml:"sISOCurrencyCode"`
}

type CountryFlagResponse struct {
	CountryFlagResult string `xml:"CountryFlagResult"`
}

type CountryFlag struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ListOfCurrenciesByName struct {
}

type ListOfCountryNamesGroupedByContinent struct {
}

type tContinent struct {
	SCode string `xml:"sCode"`
	SName string `xml:"sName"`
}

type ArrayOftCountryCodeAndName struct {
	TCountryCodeAndName []tCountryCodeAndName `xml:"tCountryCodeAndName"`
}

type ArrayOftCurrency struct {
	TCurrency []tCurrency `xml:"tCurrency"`
}

type tCurrency struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type CountryISOCode struct {
	SCountryName string `xml:"sCountryName"`
}

type tCountryCodeAndNameGroupedByContinent struct {
	Continent tContinent `xml:"Continent"`
	CountryCodeAndNames ArrayOftCountryCodeAndName `xml:"CountryCodeAndNames"`
}

type ListOfCountryNamesGroupedByContinentResponse struct {
	ListOfCountryNamesGroupedByContinentResult ArrayOftCountryCodeAndNameGroupedByContinent `xml:"ListOfCountryNamesGroupedByContinentResult"`
}

type LanguageISOCode struct {
	SLanguageName string `xml:"sLanguageName"`
}

type CapitalCityResponse struct {
	CapitalCityResult string `xml:"CapitalCityResult"`
}

type CurrencyName struct {
	SCurrencyISOCode string `xml:"sCurrencyISOCode"`
}

type ListOfCountryNamesByCodeResponse struct {
	ListOfCountryNamesByCodeResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByCodeResult"`
}

type FullCountryInfo struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type FullCountryInfoAllCountriesResponse struct {
	FullCountryInfoAllCountriesResult ArrayOftCountryInfo `xml:"FullCountryInfoAllCountriesResult"`
}

type ListOfLanguagesByCodeResponse struct {
	ListOfLanguagesByCodeResult ArrayOftLanguage `xml:"ListOfLanguagesByCodeResult"`
}

type ListOfCountryNamesByName struct {
}

type CountryCurrencyResponse struct {
	CountryCurrencyResult tCurrency `xml:"CountryCurrencyResult"`
}

type ListOfCountryNamesByNameResponse struct {
	ListOfCountryNamesByNameResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByNameResult"`
}

type ListOfCountryNamesByCode struct {
}

type ListOfCurrenciesByCode struct {
}

type tCountryCodeAndName struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type tLanguage struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
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

type CountryIntPhoneCode struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ArrayOftCountryInfo struct {
	TCountryInfo []tCountryInfo `xml:"tCountryInfo"`
}

type ListOfLanguagesByName struct {
}

type LanguageNameResponse struct {
	LanguageNameResult string `xml:"LanguageNameResult"`
}

type LanguageName struct {
	SISOCode string `xml:"sISOCode"`
}

type ArrayOftContinent struct {
	TContinent []tContinent `xml:"tContinent"`
}

type ListOfCurrenciesByCodeResponse struct {
	ListOfCurrenciesByCodeResult ArrayOftCurrency `xml:"ListOfCurrenciesByCodeResult"`
}

type ListOfContinentsByNameResponse struct {
	ListOfContinentsByNameResult ArrayOftContinent `xml:"ListOfContinentsByNameResult"`
}

type ArrayOftLanguage struct {
	TLanguage []tLanguage `xml:"tLanguage"`
}

type ListOfContinentsByName struct {
}

type FullCountryInfoResponse struct {
	FullCountryInfoResult tCountryInfo `xml:"FullCountryInfoResult"`
}

type CountryName struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type CountriesUsingCurrencyResponse struct {
	CountriesUsingCurrencyResult ArrayOftCountryCodeAndName `xml:"CountriesUsingCurrencyResult"`
}

type ListOfLanguagesByNameResponse struct {
	ListOfLanguagesByNameResult ArrayOftLanguage `xml:"ListOfLanguagesByNameResult"`
}

type LanguageISOCodeResponse struct {
	LanguageISOCodeResult string `xml:"LanguageISOCodeResult"`
}

type CountryISOCodeResponse struct {
	CountryISOCodeResult string `xml:"CountryISOCodeResult"`
}

type ListOfCurrenciesByNameResponse struct {
	ListOfCurrenciesByNameResult ArrayOftCurrency `xml:"ListOfCurrenciesByNameResult"`
}

type CountryIntPhoneCodeResponse struct {
	CountryIntPhoneCodeResult string `xml:"CountryIntPhoneCodeResult"`
}

type CurrencyNameResponse struct {
	CurrencyNameResult string `xml:"CurrencyNameResult"`
}

type CapitalCity struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type CountryCurrency struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type FullCountryInfoAllCountries struct {
}

type ListOfLanguagesByCode struct {
}

type ArrayOftCountryCodeAndNameGroupedByContinent struct {
	TCountryCodeAndNameGroupedByContinent []tCountryCodeAndNameGroupedByContinent `xml:"tCountryCodeAndNameGroupedByContinent"`
}

type ListOfContinentsByCode struct {
}

type ListOfContinentsByCodeResponse struct {
	ListOfContinentsByCodeResult ArrayOftContinent `xml:"ListOfContinentsByCodeResult"`
}

type CountryNameResponse struct {
	CountryNameResult string `xml:"CountryNameResult"`
}



// CountryFlagParser parses the SOAP response for the CountryFlag operation
func CountryFlagParse(xmlData []byte) (string, error) {
	fmt.Println(string(xmlData))

	// Define the SOAP envelope structure with the proper response type
	var response struct {
		XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Body    struct {
			Response CountryFlagResponse `xml:"CountryFlagResponse"`
		} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	}

	// Unmarshal the XML into our strongly-typed struct
	if err := xml.Unmarshal(xmlData, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("config/templates/response.tmpl")
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, response.Body.Response); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	return buf.String(), nil
}
