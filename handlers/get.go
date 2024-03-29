package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	v := fmt.Sprintf("%v", r.Context().Value("valid"))
	cd := r.Context().Value("header")

	rw.Header().Add("Content-Type", "application/json")

	var response any

	isValid, err := strconv.ParseBool(v)
	if err != nil {
		c.l.Println("[ERROR] parsing boolean", err)
		rw.WriteHeader(http.StatusInternalServerError)
		response = data.ApiError{"internal error"}
	}

	if isValid {
		c.l.Println("[DEBUG] getting country data", cd)
		rw.WriteHeader(http.StatusOK)
		response = cd
	} else {
		rw.Header().Add("Content-Type", "application/json")
		params := mux.Vars(r)
		countryCode := params["id"]
		c.l.Println("[DEBUG] country ALPHA-3 not found", countryCode)
		rw.WriteHeader(http.StatusNotFound)
		response = data.ApiError{fmt.Sprintf("country ALPHA-3, %s, not found", countryCode)}
	}

	util.ToJSON(response, rw)
}
