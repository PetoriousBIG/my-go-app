package util

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/PetoriousBIG/my-go-app/domain"
)

func ReadCountryCSV(filepath string) (*map[string]domain.CountryHeader, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	records, err := readFile(f)
	if err != nil {
		return nil, err
	}

	dict, err := validateCountryDictionary(records)
	if err != nil {
		return nil, err
	}

	return &dict, nil
}

func ReadCurrencyCSV(filepath string) (*map[string]domain.CurrencyCode, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	records, err := readFile(f)
	if err != nil {
		return nil, err
	}

	dict, err := validateCurrencyDictionary(records)
	if err != nil {
		return nil, err
	}

	return &dict, nil
}

func readFile(reader io.Reader) ([][]string, error) {
	r := csv.NewReader(reader)
	r.LazyQuotes = true

	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil

}

func validateCountryDictionary(records [][]string) (map[string]domain.CountryHeader, error) {

	countries := make(map[string]domain.CountryHeader)

	numRows := len(records)
	if numRows < 1 {
		return nil, errors.New("no rows in csv found")
	}

	for _, row := range records {

		numCols := len(row)
		if numCols != 6 {
			return nil, errors.New("incorrect number of fields found in row")
		}

		name := row[0]

		id, err := strconv.Atoi(row[3])
		if err != nil {
			return nil, err
		}

		lat, err := strconv.ParseFloat(row[4], 32)
		if err != nil {
			return nil, err
		}

		long, err := strconv.ParseFloat(row[5], 32)
		if err != nil {
			return nil, err
		}

		alpha3Code := row[2]
		alpha2Code := row[1]
		countries[alpha3Code] = domain.CountryHeader{alpha3Code, name, alpha2Code, id, lat, long}

	}

	return countries, nil
}

func validateCurrencyDictionary(records [][]string) (map[string]domain.CurrencyCode, error) {
	currencyCodes := make(map[string]domain.CurrencyCode)

	numRows := len(records)
	if numRows < 1 {
		return nil, errors.New("no rows in csv found")
	}

	for _, row := range records {

		numCols := len(row)
		if numCols != 4 {
			return nil, errors.New("incorrect number of fields found in row")
		}

		name := row[0]
		alpha2Code := row[1]
		currencyName := row[2]
		currencyCode := row[3]
		currencyCodes[alpha2Code] = domain.CurrencyCode{name, alpha2Code, currencyName, currencyCode}

	}

	return currencyCodes, nil
}
