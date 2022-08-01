package util

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/PetoriousBIG/my-go-app/data"
)

func ReadCountryCSV(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	records, err := readFile(f)
	if err != nil {
		return err
	}

	err = writeCountryDictionary(records)
	if err != nil {
		return err
	}

	return nil
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

func writeCountryDictionary(records [][]string) error {

	numRows := len(records)
	if numRows < 1 {
		return errors.New("no rows in csv found")
	}

	for _, row := range records {

		numCols := len(row)
		if numCols != 6 {
			return errors.New("incorrect number of fields found in row")
		}

		name := row[0]

		id, err := strconv.Atoi(row[3])
		if err != nil {
			return err
		}

		lat, err := strconv.ParseFloat(row[4], 32)
		if err != nil {
			return err
		}

		long, err := strconv.ParseFloat(row[5], 32)
		if err != nil {
			return err
		}

		alpha3Code := row[2]
		c := data.CountryHeader{alpha3Code, name, id, lat, long}
		data.CountryDictionary[alpha3Code] = c
	}

	return nil
}
