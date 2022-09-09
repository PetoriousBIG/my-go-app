package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/PetoriousBIG/my-go-app/data"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_GetMiddlewareValidateCountryFuncGoodParam(t *testing.T) {

	//setup - building a country dictionary
	chTST := data.CountryHeader{"TST", "Testland", "TS", 1, 0, 0}
	chNEW := data.CountryHeader{"NEW", "Newland", "NL", 2, 1, 1}
	headers := make(map[string]data.CountryHeader)

	headers[chTST.Key] = chTST
	headers[chNEW.Key] = chNEW

	//setup - building a currency dictionary
	curTST := data.CurrencyCode{"Testland", "TS", "Testland Dollars", "TSD"}
	curNEW := data.CurrencyCode{"Newland", "NL", "Newland Dollars", "NLD"}
	currencyCodes := make(map[string]data.CurrencyCode)
	currencyCodes[curTST.Alpha2Code] = curTST
	currencyCodes[curNEW.Alpha2Code] = curNEW

	//setup - HTTP handler func that will make assertions
	nextHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		header := r.Context().Value("header")
		valid := r.Context().Value("valid")
		assert.Equal(t, headers["TST"], header)
		assert.Equal(t, true, valid)
	})

	//setup - dependencies to calling GetMiddlewareValidateCountryFunc
	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
	c := NewCountryData(l)

	//setup - grabbing the functions we're testing
	function := c.GetMiddlewareValidateCountryFunc(headers, currencyCodes)
	out := function(nextHandler)

	//setup - creating URL parameters
	req := httptest.NewRequest("GET", "/", nil)
	vars := map[string]string{
		"id": "TST",
	}
	req = mux.SetURLVars(req, vars)

	//call nextHandler, which will perform assertions
	out.ServeHTTP(httptest.NewRecorder(), req)
}

func Test_GetMiddlewareValidateCountryFuncBadParam(t *testing.T) {

	//setup - building a country dictionary
	chNEW := data.CountryHeader{"NEW", "Newland", "NL", 2, 1, 1}
	headers := make(map[string]data.CountryHeader)
	headers[chNEW.Key] = chNEW

	//setup - building a currency dictionary
	curNEW := data.CurrencyCode{"Newland", "NL", "Newland Dollars", "NLD"}
	currencyCodes := make(map[string]data.CurrencyCode)
	currencyCodes[curNEW.Alpha2Code] = curNEW

	//setup - HTTP handler func that will make assertions
	nextHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		header := r.Context().Value("header")
		valid := r.Context().Value("valid")
		assert.Nil(t, header)
		assert.Equal(t, false, valid)
	})

	//setup - dependencies to calling GetMiddlewareValidateCountryFunc
	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
	c := NewCountryData(l)

	//setup - grabbing the functions we're testing
	function := c.GetMiddlewareValidateCountryFunc(headers, currencyCodes)
	out := function(nextHandler)

	//setup - creating URL parameters
	req := httptest.NewRequest("GET", "/", nil)
	vars := map[string]string{
		"id": "TST",
	}
	req = mux.SetURLVars(req, vars)

	//call nextHandler, which will perform assertions
	out.ServeHTTP(httptest.NewRecorder(), req)
}
