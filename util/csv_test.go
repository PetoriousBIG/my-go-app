package util

import (
	"bytes"
	"os"
	"testing"

	"github.com/PetoriousBIG/my-go-app/data"
	"github.com/stretchr/testify/assert"
)

func Test_ReadCountryCSVNoError(t *testing.T) {
	expected := [][]string{
		{"Peterland", "PL", "PTL", "0", "0", "0"},
	}
	var buffer bytes.Buffer
	buffer.WriteString("Peterland,PL,PTL,0,0,0")

	content, err := readFile(&buffer)

	assert.Nil(t, err)
	assert.Equal(t, expected, content)

}

func Test_ReadCountryCSVError(t *testing.T) {
	f, _ := os.Open("doesn't exist")

	content, err := readFile(f)

	assert.NotNil(t, err)
	assert.Nil(t, content)
}

func Test_WriteCountryDictionaryNoError(t *testing.T) {
	expected := make(map[string]data.CountryHeader)
	expected["PTL"] = data.CountryHeader{"PTL", "Peterland", "PL", 0, 0, 0}
	expected["CTL"] = data.CountryHeader{"CTL", "Cataland", "CL", 9, 9, 9}
	given := [][]string{
		{"Peterland", "PL", "PTL", "0", "0", "0"},
		{"Cataland", "CL", "CTL", "9", "9", "9"},
	}

	cd, err := validateCountryDictionary(given)

	assert.Nil(t, err)
	assert.Equal(t, expected, cd)
}

func Test_WriteCountryDictionaryIndexOutOfBoundsError(t *testing.T) {
	given := [][]string{
		{"Not enough fields"},
		{"Not enough fields 2"},
	}
	cd, err := validateCountryDictionary(given)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "incorrect number of fields found in row")
	assert.Nil(t, cd)
}

func Test_WriteCountryDictionaryNoRowsError(t *testing.T) {
	given := [][]string{}
	cd, err := validateCountryDictionary(given)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "no rows in csv found")
	assert.Nil(t, cd)
}

func Test_WriteCountryDictionaryIntegerConversionError(t *testing.T) {
	given := [][]string{
		{"", "", "", "invalid", "", ""},
	}
	expected := "strconv.Atoi: parsing \"invalid\": invalid syntax"
	cd, err := validateCountryDictionary(given)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expected)
	assert.Nil(t, cd)
}

func Test_WriteCountryDictionaryFloatConversionError(t *testing.T) {
	given := [][]string{
		{"", "", "", "0", "invalid", "0"},
	}
	expected := "strconv.ParseFloat: parsing \"invalid\": invalid syntax"
	cd, err := validateCountryDictionary(given)
	assert.NotNil(t, err)
	assert.EqualError(t, err, expected)
	assert.Nil(t, cd)
}
