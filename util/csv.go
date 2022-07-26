package util

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/PetoriousBIG/docker-ex/data"
)

func ReadCountryCSV(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.LazyQuotes = true

	_, err = csvReader.Read() // skip first line
	if err != nil {
		return err
	}

	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	err = writeCountryDictionary(records)
	if err != nil {
		return err
	}

	return nil
}

func writeCountryDictionary(records [][]string) error {

	for _, row := range records {

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
