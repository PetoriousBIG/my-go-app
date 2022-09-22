package domain

// Country defines the structure for an API object
type AtAGlance struct {
	CountryHeader CountryHeader `json:"country_header"`
	ExchangeRates Finance       `json:"exchange_rates"`
}

type CountryHeader struct {
	Key              string  `json:"country"`
	Name             string  `json:"country_name"`
	Alpha2Code       string  `json:"alpha2"`
	Id               int     `json:"country_id"`
	AverageLatitude  float64 `json:"average_latitude"`
	AverageLongitude float64 `json:"average_longitude"`
}
