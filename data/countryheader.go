package data

type CountryHeader struct {
	Key              string  `json:"country"`
	Name             string  `json:"country_name"`
	Id               int     `json:"country_id"`
	AverageLatitude  float64 `json:"average_latitude"`
	AverageLongitude float64 `json:"average_longitude"`
}
