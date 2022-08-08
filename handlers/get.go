package handlers

import (
	"log"
	"net/http"

	"github.com/PetoriousBIG/my-go-app/util"
)

type countryData struct {
	l *log.Logger
}

func NewCountryData(l *log.Logger) *countryData {
	return &countryData{l}
}

func (c *countryData) GetCountryData(rw http.ResponseWriter, r *http.Request) {
	header := r.Context().Value("header")
	c.l.Println("[DEBUG] Get Country Data", header)

	err := util.ToJSON(header, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
