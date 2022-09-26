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
	clients.ClientStruct = &getClientMock{} // use the mock

	response := FinanceProvider.GetFinance(domain.FinanceRequest{"USD", "abc_123"})
	assert.NotNil(t, response)
	assert.EqualValues(t, "USD", response.Base)
	assert.EqualValues(t, 123.55, response.Rates["EUR"])
	assert.EqualValues(t, 0.0054, response.Rates["JPY"])
	assert.EqualValues(t, true, response.Success)
	assert.EqualValues(t, "2022-09-22", response.Date)
	assert.EqualValues(t, 0, response.Error.Code)
	assert.EqualValues(t, "", response.Error.Message)
}

func TestGetFinanceInvalidBase(t *testing.T) {
	getRequestFunc = func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{"base":"", "rates":{},"success":false, "date":"", "error":{"code":201, "type":"invalid_base_currency"}}`)),
		}, nil
	}
	clients.ClientStruct = &getClientMock{}

	response := FinanceProvider.GetFinance(domain.FinanceRequest{"AAA", ""})

	assert.NotNil(t, response)
	assert.EqualValues(t, "", response.Base)
	assert.Empty(t, response.Rates)
	assert.EqualValues(t, false, response.Success)
	assert.EqualValues(t, "", response.Date)
	assert.EqualValues(t, 201, response.Error.Code)
	assert.EqualValues(t, "invalid_base_currency", response.Error.Message)

}
