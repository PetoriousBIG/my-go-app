package data

import (
	"encoding/json"
	"io"
)

type CountryHeader struct {
	Key              string  `json:"country"`
	Name             string  `json:"country_name"`
	Id               int     `json:"country_id"`
	AverageLatitude  float64 `json:"average_latitude"`
	AverageLongitude float64 `json:"average_longitude"`
}

var CountryDictionary = make(map[string]CountryHeader)

// ToJSON serializes the given interface into a string based JSON format
func (c *CountryHeader) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(c)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func (c *CountryHeader) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c)
}
