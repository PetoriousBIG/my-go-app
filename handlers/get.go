package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PetoriousBIG/my-go-app/clients"
	"github.com/PetoriousBIG/my-go-app/data"
	"github.com/PetoriousBIG/my-go-app/util"
	"github.com/gorilla/mux"
)

type countryData struct {
	l *log.Logger
}

func NewCountryData(l *log.Logger) *countryData {
	return &countryData{l}
}

func (c *countryData) GetCountryData(rw http.ResponseWriter, r *http.Request) {
	valid := r.Context().Value("valid").(bool)
	header := r.Context().Value("header").(data.CountryHeader)

	rw.Header().Add("Content-Type", "application/json")

	var response any

	if valid {
		c.l.Println("[DEBUG] getting country data", header)
		rw.WriteHeader(http.StatusOK)
		currency := r.Context().Value("currency").(data.CurrencyCode)
		response = apiMashup(header, currency)
	} else {
		c.l.Println("[DEBUG] country not found", header)
		rw.WriteHeader(http.StatusNotFound)
		response = data.ApiError{fmt.Sprintf("country not found, %v", header)}

	}

	util.ToJSON(response, rw)
}

func apiMashup(ch data.CountryHeader, cc data.CurrencyCode) *data.AtAGlance {

	base := cc.CurrencyCode
	finance := clients.NewFinacne(base)
	exchangeRates := finance.GET()

	output := data.AtAGlance{ch, exchangeRates}

	return &output
}
