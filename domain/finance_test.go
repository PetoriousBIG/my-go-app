package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinance(t *testing.T) {
	request := Finance{
		Base: "USD",
		Rates: map[string]float64{
			"EUR": 1.000,
			"JPY": 1234.1111,
			"COP": 4400.1234532,
		},
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var result Finance
	err = json.Unmarshal(bytes, &result)

	assert.Nil(t, err)
	assert.EqualValues(t, result.Base, request.Base)
	assert.EqualValues(t, result.Rates["EUR"], request.Rates["EUR"])
	assert.EqualValues(t, result.Rates["JPY"], request.Rates["JPY"])
	assert.EqualValues(t, result.Rates["COP"], request.Rates["COP"])
}

func TestFinanceError(t *testing.T) {
	request := NewFinanceError(400, "Bad Request Error")

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var errResult financeError
	err = json.Unmarshal(bytes, &errResult)
	assert.Nil(t, err)
	assert.EqualValues(t, errResult.Status(), request.Status())
	assert.EqualValues(t, errResult.Message(), errResult.Message())
}
