package data

type CountryHeader struct {
	Key              string
	Name             string
	Id               int
	AverageLatitude  float64
	AverageLongitude float64
}

var CountryDictionary = make(map[string]CountryHeader)
