package generated

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

type CurrencyNameResponse struct {
	CurrencyNameResult string `xml:"CurrencyNameResult"`
}

type FullCountryInfo struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ListOfLanguagesByCodeResponse struct {
	ListOfLanguagesByCodeResult ArrayOftLanguage `xml:"ListOfLanguagesByCodeResult"`
}

type ListOfCountryNamesByCodeResponse struct {
	ListOfCountryNamesByCodeResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByCodeResult"`
}

type LanguageISOCode struct {
	SLanguageName string `xml:"sLanguageName"`
}

type ArrayOftCountryInfo struct {
	TCountryInfo []tCountryInfo `xml:"tCountryInfo"`
}

type CountryIntPhoneCode struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type FullCountryInfoAllCountriesResponse struct {
	FullCountryInfoAllCountriesResult ArrayOftCountryInfo `xml:"FullCountryInfoAllCountriesResult"`
}

type ArrayOftCountryCodeAndNameGroupedByContinent struct {
	TCountryCodeAndNameGroupedByContinent []tCountryCodeAndNameGroupedByContinent `xml:"tCountryCodeAndNameGroupedByContinent"`
}

type tCurrency struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type ArrayOftContinent struct {
	TContinent []tContinent `xml:"tContinent"`
}

type tLanguage struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type ListOfLanguagesByNameResponse struct {
	ListOfLanguagesByNameResult ArrayOftLanguage `xml:"ListOfLanguagesByNameResult"`
}

type CountryCurrency struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type tCountryCodeAndNameGroupedByContinent struct {
	Continent tContinent `xml:"Continent"`
	CountryCodeAndNames ArrayOftCountryCodeAndName `xml:"CountryCodeAndNames"`
}

type ListOfCountryNamesGroupedByContinent struct {
}

type ListOfLanguagesByName struct {
}

type ListOfCountryNamesByNameResponse struct {
	ListOfCountryNamesByNameResult ArrayOftCountryCodeAndName `xml:"ListOfCountryNamesByNameResult"`
}

type ArrayOftCurrency struct {
	TCurrency []tCurrency `xml:"tCurrency"`
}

type FullCountryInfoResponse struct {
	FullCountryInfoResult tCountryInfo `xml:"FullCountryInfoResult"`
}

type LanguageNameResponse struct {
	LanguageNameResult string `xml:"LanguageNameResult"`
}

type ListOfContinentsByCodeResponse struct {
	ListOfContinentsByCodeResult ArrayOftContinent `xml:"ListOfContinentsByCodeResult"`
}

type ListOfCountryNamesGroupedByContinentResponse struct {
	ListOfCountryNamesGroupedByContinentResult ArrayOftCountryCodeAndNameGroupedByContinent `xml:"ListOfCountryNamesGroupedByContinentResult"`
}

type ListOfCurrenciesByNameResponse struct {
	ListOfCurrenciesByNameResult ArrayOftCurrency `xml:"ListOfCurrenciesByNameResult"`
}

type ArrayOftCountryCodeAndName struct {
	TCountryCodeAndName []tCountryCodeAndName `xml:"tCountryCodeAndName"`
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

type CountriesUsingCurrencyResponse struct {
	CountriesUsingCurrencyResult ArrayOftCountryCodeAndName `xml:"CountriesUsingCurrencyResult"`
}

type CountriesUsingCurrency struct {
	SISOCurrencyCode string `xml:"sISOCurrencyCode"`
}

type CurrencyName struct {
	SCurrencyISOCode string `xml:"sCurrencyISOCode"`
}

type CapitalCityResponse struct {
	CapitalCityResult string `xml:"CapitalCityResult"`
}

type ListOfCurrenciesByName struct {
}

type CountryISOCodeResponse struct {
	CountryISOCodeResult string `xml:"CountryISOCodeResult"`
}

type CountryIntPhoneCodeResponse struct {
	CountryIntPhoneCodeResult string `xml:"CountryIntPhoneCodeResult"`
}

type ListOfCountryNamesByCode struct {
}

type FullCountryInfoAllCountries struct {
}

type CountryNameResponse struct {
	CountryNameResult string `xml:"CountryNameResult"`
}

type CapitalCity struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ListOfCountryNamesByName struct {
}

type LanguageName struct {
	SISOCode string `xml:"sISOCode"`
}

type tCountryCodeAndName struct {
	SISOCode string `xml:"sISOCode"`
	SName string `xml:"sName"`
}

type ArrayOftLanguage struct {
	TLanguage []tLanguage `xml:"tLanguage"`
}

type ListOfContinentsByCode struct {
}

type CountryISOCode struct {
	SCountryName string `xml:"sCountryName"`
}

type ListOfLanguagesByCode struct {
}

type LanguageISOCodeResponse struct {
	LanguageISOCodeResult string `xml:"LanguageISOCodeResult"`
}

type ListOfContinentsByName struct {
}

type CountryFlagResponse struct {
	CountryFlagResult string `xml:"CountryFlagResult"`
}

type CountryName struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type ListOfContinentsByNameResponse struct {
	ListOfContinentsByNameResult ArrayOftContinent `xml:"ListOfContinentsByNameResult"`
}

type ListOfCurrenciesByCode struct {
}

type ListOfCurrenciesByCodeResponse struct {
	ListOfCurrenciesByCodeResult ArrayOftCurrency `xml:"ListOfCurrenciesByCodeResult"`
}

type CountryCurrencyResponse struct {
	CountryCurrencyResult tCurrency `xml:"CountryCurrencyResult"`
}

type CountryFlag struct {
	SCountryISOCode string `xml:"sCountryISOCode"`
}

type tContinent struct {
	SCode string `xml:"sCode"`
	SName string `xml:"sName"`
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
