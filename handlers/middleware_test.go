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

// basic test to make sure the middleware is correctly utilizing the
// country dictionary to verify URL params.
func Test_GetMiddlewareValidateCountryFuncNoError(t *testing.T) {

	//setup - building a country dictionary
	chTST := data.CountryHeader{"TST", "Testland", 1, 0, 0}
	chNEW := data.CountryHeader{"NEW", "Newland", 2, 1, 1}
	testmap := make(map[string]data.CountryHeader)
	testmap[chTST.Key] = chTST
	testmap[chNEW.Key] = chNEW
	cd := data.CountryDictionary{testmap}

	//setup - HTTP handler func that will make assertions
	nextHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		header := r.Context().Value("header")
		valid := r.Context().Value("valid")
		assert.Equal(t, cd.Dict["TST"], header)
		assert.Equal(t, "yes", valid)
	})

	//setup - dependencies to calling GetMiddlewareValidateCountryFunc
	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
	c := NewCountryData(l)

	//setup - passing
	function := c.GetMiddlewareValidateCountryFunc(&cd)
	out := function(nextHandler)

	req := httptest.NewRequest("GET", "/", nil)
	vars := map[string]string{
		"id": "TST",
	}
	req = mux.SetURLVars(req, vars)

	out.ServeHTTP(httptest.NewRecorder(), req)
}

func Test_GetMiddlewareValidateCountryFuncNotFoundError(t *testing.T) {

	//setup - building a country dictionary
	chNEW := data.CountryHeader{"NEW", "Newland", 2, 1, 1}
	testmap := make(map[string]data.CountryHeader)
	testmap[chNEW.Key] = chNEW
	cd := data.CountryDictionary{testmap}

	//setup - HTTP handler func that will make assertions
	nextHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		header := r.Context().Value("header")
		valid := r.Context().Value("valid")
		assert.Nil(t, header)
		assert.Equal(t, "no", valid)
	})

	//setup - dependencies to calling GetMiddlewareValidateCountryFunc
	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
	c := NewCountryData(l)

	//setup - passing
	function := c.GetMiddlewareValidateCountryFunc(&cd)
	out := function(nextHandler)

	req := httptest.NewRequest("GET", "/", nil)
	vars := map[string]string{
		"id": "TST",
	}
	req = mux.SetURLVars(req, vars)

	out.ServeHTTP(httptest.NewRecorder(), req)
}
