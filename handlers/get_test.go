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

	req := httptest.NewRequest("GET", "/", nil)
	ctx := req.Context()
	ctx = context.WithValue(ctx, "valid", true)
	ctx = context.WithValue(ctx, "header", "test")

	out := httptest.NewRecorder()
	c.GetCountryData(out, req.WithContext(ctx))

	actualStatus := out.Result().StatusCode
	actualHeader := string(out.Body.Bytes())
	assert.Equal(t, http.StatusOK, actualStatus)
	assert.Equal(t, "test", actualHeader)

}

func Test_GetCountryDataWithInvalidCountry(t *testing.T) {
	l := log.New(os.Stdout, "UNIT TEST ", log.LstdFlags)
	c := NewCountryData(l)
	expectedHeader := data.ApiError{"country not found, test"}

	req := httptest.NewRequest("GET", "/", nil)
	ctx := req.Context()
	ctx = context.WithValue(ctx, "valid", false)
	ctx = context.WithValue(ctx, "header", "test")

	out := httptest.NewRecorder()
	c.GetCountryData(out, req.WithContext(ctx))

	actualStatus := out.Result().StatusCode
	var actualHeader data.ApiError
	util.FromJSON(&actualHeader, out.Body)

	assert.Equal(t, http.StatusNotFound, actualStatus)
	assert.Equal(t, expectedHeader, actualHeader)
}
