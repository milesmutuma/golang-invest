package my_csv

import (
	"encoding/csv"
	"opg-analysis/internal/raw"
	"os"
	"slices"
	"strconv"
)

type columns = []string

type rows = []columns

type loader struct {
	path string
}

func (l *loader) Load() ([]raw.Stock, error) {
	rows, err := l.read()
	if err != nil {
		return nil, err
	}
	var data []raw.Stock

	for _, row := range rows {
		ticker := row[0]
		gap, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			continue
		}
		openingPrice, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			continue
		}
		data = append(data, raw.Stock{
			Ticker:       ticker,
			Gap:          gap,
			OpeningPrice: openingPrice,
		})

	}
	return data, nil
}

func (l *loader) read() (rows, error) {
	f, err := os.Open(l.path)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Remove the header row
	//Equiavalent to:
	rows = slices.Delete(rows, 0, 1)
	return rows, nil
}

func NewLoader(path string) raw.Loader {
	return &loader{
		path,
	}
}
