package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/PetoriousBIG/my-go-app/data"
	"github.com/PetoriousBIG/my-go-app/util"
	"github.com/stretchr/testify/assert"
)

func Test_GetCountryDataWithValidCountry(t *testing.T) {
	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
	c := NewCountryData(l)
	expectedAtAGlance := data.AtAGlance{data.CountryHeader{"TST", "Testland", "TS", 1, 0, 0}, data.Finance{}}

	req := httptest.NewRequest("GET", "/", nil)
	ctx := req.Context()
	ctx = context.WithValue(ctx, "valid", true)
	ctx = context.WithValue(ctx, "header", expectedAtAGlance.CountryHeader)
	ctx = context.WithValue(ctx, "currency", data.CurrencyCode{"Testland", "TS", "Testland Dollars", "TSD"})

	out := httptest.NewRecorder()
	c.GetCountryData(out, req.WithContext(ctx))

	actualStatus := out.Result().StatusCode
	var actualAtAGlance data.AtAGlance
	util.FromJSON(&actualAtAGlance, out.Body)

	assert.Equal(t, http.StatusOK, actualStatus)
	assert.Equal(t, expectedAtAGlance, actualAtAGlance)

}

func Test_GetCountryDataWithInvalidCountry(t *testing.T) {
	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
	c := NewCountryData(l)
	arg := data.CountryHeader{"TST", "Testland", "TS", 1, 0, 0}
	expectedHeader := data.ApiError{"country not found, {TST Testland TS 1 0 0}"}

	req := httptest.NewRequest("GET", "/", nil)
	ctx := req.Context()
	ctx = context.WithValue(ctx, "valid", false)
	ctx = context.WithValue(ctx, "header", arg)
	ctx = context.WithValue(ctx, "currency", nil)

	out := httptest.NewRecorder()
	c.GetCountryData(out, req.WithContext(ctx))

	actualStatus := out.Result().StatusCode
	var actualHeader data.ApiError
	util.FromJSON(&actualHeader, out.Body)

	assert.Equal(t, http.StatusNotFound, actualStatus)
	assert.Equal(t, expectedHeader, actualHeader)
}
