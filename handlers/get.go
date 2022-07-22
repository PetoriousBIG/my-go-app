package handlers

import (
	"log"
	"net/http"

	"github.com/PetoriousBIG/docker-ex/data"
)

type countryData struct {
	l *log.Logger
}

func NewCountryData(l *log.Logger) *countryData {
	return &countryData{l}
}

func (c *countryData) GetCountryData(rw http.ResponseWriter, r *http.Request) {
	c.l.Println("[DEBUG] Get Country Data")

	cd := data.GetCountryData()

	err := cd.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
