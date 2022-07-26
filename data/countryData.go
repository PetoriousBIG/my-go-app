package data

import (
	"encoding/json"
	"io"
)

// Country defines the structure for an API object
type Country struct {
	CountryName string `json:"country"`
}

// ToJSON serializes the given interface into a string based JSON format
func (c *Country) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(c)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func (c *Country) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c)
}

var sample = Country{CountryName: "Colombia"}

func GetCountryData() *Country {
	return &sample
}
