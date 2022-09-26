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
		Success: true,
		Date:    "2022-09-25",
		Error: &FinanceError{
			Code:    200,
			Message: "all, good",
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
	assert.EqualValues(t, result.Success, request.Success)
	assert.EqualValues(t, result.Date, request.Date)
	assert.EqualValues(t, result.Error.Code, request.Error.Code)
	assert.EqualValues(t, result.Error.Message, request.Error.Message)
}
