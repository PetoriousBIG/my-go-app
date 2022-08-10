package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	v := fmt.Sprintf("%v", r.Context().Value("valid"))
	cd := r.Context().Value("header")
	c.l.Println("[DEBUG] Getting country data", cd)

	rw.Header().Add("Content-Type", "application/json")

	var response data.ApiError

	isValid, err := strconv.ParseBool(v)
	if err != nil {
		c.l.Println("[ERROR] parsing boolean", err)
		rw.WriteHeader(http.StatusInternalServerError)
		response = data.ApiError{"internal error"}
	}

	if isValid {

	} else {
		c.l.Println("[DEBUG] country not found", cd)
		rw.WriteHeader(http.StatusNotFound)
		response = data.ApiError{fmt.Sprintf("country not found, %v", cd)}
	}

	util.ToJSON(response, rw)
}
