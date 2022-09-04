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

var chTST = data.CountryHeader{"TST", "Testland", 1, 0, 0}
var chNEW = data.CountryHeader{"NEW", "Newland", 2, 1, 1}
var empty = data.CountryHeader{"", "", 0, 0, 0}

func Test_GetMiddlewareValidateCountryFuncGoodParam(t *testing.T) {

	testmap := make(map[string]data.CountryHeader)
	testmap[chTST.Key] = chTST
	testmap[chNEW.Key] = chNEW
	cd := data.CountryDictionary{testmap}

	//setup - HTTP handler func that will make assertions
	nextHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		header := r.Context().Value("header")
		valid := r.Context().Value("valid")
		assert.Equal(t, cd.Dict["TST"], header)
		assert.Equal(t, true, valid)
	})

	//setup - dependencies to calling GetMiddlewareValidateCountryFunc
	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
	c := NewCountryData(l)

	//setup - grabbing the functions we're testing
	function := c.GetMiddlewareValidateCountryFunc(&cd)
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

	testmap := make(map[string]data.CountryHeader)
	testmap[chNEW.Key] = chNEW
	cd := data.CountryDictionary{testmap}

	//setup - HTTP handler func that will make assertions
	nextHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		header := r.Context().Value("header")
		valid := r.Context().Value("valid")
		assert.Equal(t, empty, header)
		assert.Equal(t, false, valid)
	})

	//setup - dependencies to calling GetMiddlewareValidateCountryFunc
	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
	c := NewCountryData(l)

	//setup - grabbing the functions we're testing
	function := c.GetMiddlewareValidateCountryFunc(&cd)
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
