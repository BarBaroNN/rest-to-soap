package generated

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

type FullCountryInfoResponse struct {
	FullCountryInfoResult tCountryInfo `xml:"FullCountryInfoResult"`
}

type LanguageNameResponse struct {
	LanguageNameResult string `xml:"LanguageNameResult"`
}

type CountryName struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ArrayOftLanguage struct {
	TLanguage []tLanguage `xml:"tLanguage"`
}

type ArrayOftContinent struct {
	TContinent []tContinent `xml:"tContinent"`
}

type ListOfCurrenciesByName struct {
}

type CountryIntPhoneCodeResponse struct {
	CountryIntPhoneCodeResult string `xml:"CountryIntPhoneCodeResult"`
}

type ListOfLanguagesByName struct {
}

type LanguageISOCodeResponse struct {
	LanguageISOCodeResult string `xml:"LanguageISOCodeResult"`
}

type ListOfCountryNamesGroupedByContinent struct {
}

type ListOfLanguagesByCode struct {
}

type ArrayOftCurrency struct {
	TCurrency []tCurrency `xml:"tCurrency"`
}

type FullCountryInfoAllCountriesResponse struct {
	FullCountryInfoAllCountriesResult ArrayOftCountryInfo `xml:"FullCountryInfoAllCountriesResult"`
}

type CurrencyNameResponse struct {
	CurrencyNameResult string `xml:"CurrencyNameResult"`
}

type CountryCurrencyResponse struct {
	CountryCurrencyResult tCurrency `xml:"CountryCurrencyResult"`
}

type CurrencyName struct {
	SCurrencyISOCode string `xml:"sCurrencyISOCode"`
}

type ListOfLanguagesByNameResponse struct {
	ListOfLanguagesByNameResult ArrayOftLanguage `xml:"ListOfLanguagesByNameResult"`
}

type CountryFlag struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type CountriesUsingCurrencyResponse struct {
	CountriesUsingCurrencyResult ArrayOftCountryCodeAndName `xml:"CountriesUsingCurrencyResult"`
}

type CountryISOCode struct {
	SCountryName string `xml:"sCountryName"`
}

type CountryIntPhoneCode struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ArrayOftCountryCodeAndName struct {
	TCountryCodeAndName []tCountryCodeAndName `xml:"tCountryCodeAndName"`
}

type ListOfContinentsByNameResponse struct {
	ListOfContinentsByNameResult ArrayOftContinent `xml:"ListOfContinentsByNameResult"`
}

type ListOfCurrenciesByCodeResponse struct {
	ListOfCurrenciesByCodeResult ArrayOftCurrency `xml:"ListOfCurrenciesByCodeResult"`
}

type CountryCurrency struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type FullCountryInfoAllCountries struct {
}

type CountryFlagResponse struct {
	CountryFlagResult string `xml:"CountryFlagResult"`
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

type ListOfContinentsByCodeResponse struct {
	ListOfContinentsByCodeResult ArrayOftContinent `xml:"ListOfContinentsByCodeResult"`
}

type ListOfCountryNamesByCode struct {
}

type ListOfCountryNamesByNameResponse struct {
	ListOfCountryNamesByNameResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByNameResult"`
}

type ArrayOftCountryCodeAndNameGroupedByContinent struct {
	TCountryCodeAndNameGroupedByContinent []tCountryCodeAndNameGroupedByContinent `xml:"tCountryCodeAndNameGroupedByContinent"`
}

type LanguageName struct {
	SISOCode string `xml:"sISOCode"`
}

type ListOfContinentsByCode struct {
}

type tContinent struct {
	SCode string `xml:"sCode"`
	SName string `xml:"sName"`
}

type tCountryCodeAndName struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type ListOfCountryNamesByName struct {
}

type ListOfCurrenciesByCode struct {
}

type CountriesUsingCurrency struct {
	SISOCurrencyCode string `xml:"sISOCurrencyCode"`
}

type ListOfLanguagesByCodeResponse struct {
	ListOfLanguagesByCodeResult ArrayOftLanguage `xml:"ListOfLanguagesByCodeResult"`
}

type FullCountryInfo struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ListOfCountryNamesByCodeResponse struct {
	ListOfCountryNamesByCodeResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByCodeResult"`
}

type CountryNameResponse struct {
	CountryNameResult string `xml:"CountryNameResult"`
}

type tLanguage struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type ArrayOftCountryInfo struct {
	TCountryInfo []tCountryInfo `xml:"tCountryInfo"`
}

type CapitalCityResponse struct {
	CapitalCityResult string `xml:"CapitalCityResult"`
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

type ListOfCurrenciesByNameResponse struct {
	ListOfCurrenciesByNameResult ArrayOftCurrency `xml:"ListOfCurrenciesByNameResult"`
}

type CountryISOCodeResponse struct {
	CountryISOCodeResult string `xml:"CountryISOCodeResult"`
}

type tCurrency struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type ListOfContinentsByName struct {
}

type CapitalCity struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
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
