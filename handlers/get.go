package handlers

import (
	"log"
	"net/http"

	"github.com/PetoriousBIG/my-go-app/data"
	"github.com/gorilla/mux"
)

type countryData struct {
	l *log.Logger
}

func NewCountryData(l *log.Logger) *countryData {
	return &countryData{l}
}

func (c *countryData) GetCountryData(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	countryCode := params["id"]

	c.l.Println("[DEBUG] Get Country Data", countryCode)

	cd := data.CountryDictionary[countryCode]

	err := cd.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
