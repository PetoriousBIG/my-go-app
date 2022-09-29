package providers

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/PetoriousBIG/my-go-app/clients"
	"github.com/PetoriousBIG/my-go-app/domain"
	"github.com/stretchr/testify/assert"
)

var (
	getRequestFunc func(r *http.Request) (*http.Response, error)
)

type getClientMock struct{}

//We are mocking the client method "Get"
func (cm *getClientMock) Get(r *http.Request) (*http.Response, error) {
	return getRequestFunc(r)
}

func TestGetFinanceSuccessfulCall(t *testing.T) {
	getRequestFunc = func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"base":"USD", "rates":{"EUR":123.55, "JPY":0.0054}, "success":true, "date":"2022-09-22", "error":{}}`)),
		}, nil
	}
	clients.ClientStruct = &getClientMock{}

	response, err := ExchangeProvider.GetExchange(domain.ExchangeRequest{"USD", "abc_123"})
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, "USD", response.Base)
	assert.EqualValues(t, 123.55, response.Rates["EUR"])
	assert.EqualValues(t, 0.0054, response.Rates["JPY"])
	assert.EqualValues(t, "2022-09-22", response.Date)
}

func TestGetFinanceInvalidBase(t *testing.T) {
	getRequestFunc = func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"base":"", "rates":{},"success":false, "date":"", "error":{"code":201, "type":"invalid_base_currency"}}`)),
		}, nil
	}
	clients.ClientStruct = &getClientMock{}

	response, err := ExchangeProvider.GetExchange(domain.ExchangeRequest{"AAA", "abc_123"})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, 201, err.Code)
	assert.EqualValues(t, "invalid_base_currency", err.Message)
}

func TestGetFinanceInvalidAPIKey(t *testing.T) {
	getRequestFunc = func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"base":"", "rates":{},"success":false, "date":"", "error":{"code":null, "type":"Invalid authentication credentials"}}`)),
		}, nil
	}
	clients.ClientStruct = &getClientMock{}

	response, err := ExchangeProvider.GetExchange(domain.ExchangeRequest{"AAA", "bad_key"})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, 0, err.Code)
	assert.EqualValues(t, "Invalid authentication credentials", err.Message)
}
