package util

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/PetoriousBIG/my-go-app/data"
)

func ReadCountryCSV(filepath string) (*data.CountryDictionary, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	records, err := readFile(f)
	if err != nil {
		return nil, err
	}

	cd, err := validateCountryDictionary(records)
	if err != nil {
		return nil, err
	}

	dict := data.CountryDictionary{cd}
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

func validateCountryDictionary(records [][]string) (map[string]data.CountryHeader, error) {

	cd := make(map[string]data.CountryHeader)

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
		cd[alpha3Code] = data.CountryHeader{alpha3Code, name, id, lat, long}

	}

	return cd, nil
}
