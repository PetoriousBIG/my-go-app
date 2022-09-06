package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PetoriousBIG/my-go-app/clients"
	"github.com/PetoriousBIG/my-go-app/data"
	"github.com/PetoriousBIG/my-go-app/util"
)

type countryData struct {
	l *log.Logger
}

func NewCountryData(l *log.Logger) *countryData {
	return &countryData{l}
}

func (c *countryData) GetCountryData(rw http.ResponseWriter, r *http.Request) {
	valid := r.Context().Value("valid").(bool)
	ch := r.Context().Value("header").(data.CountryHeader)

	rw.Header().Add("Content-Type", "application/json")

	var response any

	if valid {
		c.l.Println("[DEBUG] getting country data", ch)
		rw.WriteHeader(http.StatusOK)
		response = apiMashup(ch)
	} else {
		c.l.Println("[DEBUG] country not found", ch)
		rw.WriteHeader(http.StatusNotFound)
		response = data.ApiError{fmt.Sprintf("country not found, %v", ch)}
	}

	util.ToJSON(response, rw)
}

func apiMashup(ch data.CountryHeader) any {
	base := ch.Key
	finance := clients.NewFinacne(base)
	retval := finance.GET()
	fmt.Printf("%v", retval)
	return nil
}
