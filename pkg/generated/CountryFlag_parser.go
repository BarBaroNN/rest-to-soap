package generated

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

type CountryCurrencyResponse struct {
	CountryCurrencyResult tCurrency `xml:"CountryCurrencyResult"`
}

type ListOfLanguagesByCodeResponse struct {
	ListOfLanguagesByCodeResult ArrayOftLanguage `xml:"ListOfLanguagesByCodeResult"`
}

type ArrayOftContinent struct {
	TContinent []tContinent `xml:"tContinent"`
}

type ArrayOftCountryInfo struct {
	TCountryInfo []tCountryInfo `xml:"tCountryInfo"`
}

type FullCountryInfoAllCountriesResponse struct {
	FullCountryInfoAllCountriesResult ArrayOftCountryInfo `xml:"FullCountryInfoAllCountriesResult"`
}

type ListOfCountryNamesGroupedByContinentResponse struct {
	ListOfCountryNamesGroupedByContinentResult ArrayOftCountryCodeAndNameGroupedByContinent `xml:"ListOfCountryNamesGroupedByContinentResult"`
}

type ListOfContinentsByNameResponse struct {
	ListOfContinentsByNameResult ArrayOftContinent `xml:"ListOfContinentsByNameResult"`
}

type ListOfContinentsByCodeResponse struct {
	ListOfContinentsByCodeResult ArrayOftContinent `xml:"ListOfContinentsByCodeResult"`
}

type ListOfCurrenciesByName struct {
}

type ListOfCurrenciesByNameResponse struct {
	ListOfCurrenciesByNameResult ArrayOftCurrency `xml:"ListOfCurrenciesByNameResult"`
}

type CountryFlagResponse struct {
	CountryFlagResult string `xml:"CountryFlagResult"`
}

type ListOfContinentsByCode struct {
}

type tCountryCodeAndName struct {
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

type ListOfCountryNamesByNameResponse struct {
	ListOfCountryNamesByNameResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByNameResult"`
}

type ArrayOftLanguage struct {
	TLanguage []tLanguage `xml:"tLanguage"`
}

type FullCountryInfoResponse struct {
	FullCountryInfoResult tCountryInfo `xml:"FullCountryInfoResult"`
}

type CountryISOCodeResponse struct {
	CountryISOCodeResult string `xml:"CountryISOCodeResult"`
}

type ListOfCountryNamesByName struct {
}

type CountryFlag struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ListOfCurrenciesByCodeResponse struct {
	ListOfCurrenciesByCodeResult ArrayOftCurrency `xml:"ListOfCurrenciesByCodeResult"`
}

type CountryISOCode struct {
	SCountryName string `xml:"sCountryName"`
}

type ArrayOftCountryCodeAndNameGroupedByContinent struct {
	TCountryCodeAndNameGroupedByContinent []tCountryCodeAndNameGroupedByContinent `xml:"tCountryCodeAndNameGroupedByContinent"`
}

type FullCountryInfo struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ListOfCountryNamesGroupedByContinent struct {
}

type LanguageNameResponse struct {
	LanguageNameResult string `xml:"LanguageNameResult"`
}

type ListOfLanguagesByName struct {
}

type CountryIntPhoneCode struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type FullCountryInfoAllCountries struct {
}

type tCountryCodeAndNameGroupedByContinent struct {
	Continent tContinent `xml:"Continent"`
	CountryCodeAndNames ArrayOftCountryCodeAndName `xml:"CountryCodeAndNames"`
}

type tCurrency struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type LanguageISOCode struct {
	SLanguageName string `xml:"sLanguageName"`
}

type ListOfCurrenciesByCode struct {
}

type ListOfLanguagesByNameResponse struct {
	ListOfLanguagesByNameResult ArrayOftLanguage `xml:"ListOfLanguagesByNameResult"`
}

type ArrayOftCurrency struct {
	TCurrency []tCurrency `xml:"tCurrency"`
}

type CountriesUsingCurrency struct {
	SISOCurrencyCode string `xml:"sISOCurrencyCode"`
}

type tContinent struct {
	SCode string `xml:"sCode"`
	SName string `xml:"sName"`
}

type ListOfContinentsByName struct {
}

type ListOfCountryNamesByCodeResponse struct {
	ListOfCountryNamesByCodeResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByCodeResult"`
}

type CapitalCity struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type LanguageISOCodeResponse struct {
	LanguageISOCodeResult string `xml:"LanguageISOCodeResult"`
}

type CountryIntPhoneCodeResponse struct {
	CountryIntPhoneCodeResult string `xml:"CountryIntPhoneCodeResult"`
}

type ListOfCountryNamesByCode struct {
}

type LanguageName struct {
	SISOCode string `xml:"sISOCode"`
}

type CountryCurrency struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ListOfLanguagesByCode struct {
}

type CountryNameResponse struct {
	CountryNameResult string `xml:"CountryNameResult"`
}

type CountryName struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ArrayOftCountryCodeAndName struct {
	TCountryCodeAndName []tCountryCodeAndName `xml:"tCountryCodeAndName"`
}

type CountriesUsingCurrencyResponse struct {
	CountriesUsingCurrencyResult ArrayOftCountryCodeAndName `xml:"CountriesUsingCurrencyResult"`
}

type tLanguage struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type CurrencyName struct {
	SCurrencyISOCode string `xml:"sCurrencyISOCode"`
}

type CurrencyNameResponse struct {
	CurrencyNameResult string `xml:"CurrencyNameResult"`
}

type CapitalCityResponse struct {
	CapitalCityResult string `xml:"CapitalCityResult"`
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
