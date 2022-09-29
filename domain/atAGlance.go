package domain

// Country defines the structure for an API object
type AtAGlance struct {
	CountryHeader CountryHeader `json:"country_header"`
	ExchangeRates Exchange      `json:"exchange_rates"`
}

type CountryHeader struct {
	Key              string  `json:"country"`
	Name             string  `json:"country_name"`
	Alpha2Code       string  `json:"alpha2"`
	Id               int     `json:"country_id"`
	AverageLatitude  float64 `json:"average_latitude"`
	AverageLongitude float64 `json:"average_longitude"`
}

type CurrencyCode struct {
	Name         string
	Alpha2Code   string
	CurrencyName string
	CurrencyCode string
}

type AtAGlanceError struct {
}
