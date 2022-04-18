package lib

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadInputCsv(inputFile string) []string {

	csvFile, err := os.Open(inputFile)
	var csvList []string

	if err != nil {

		log.Fatal(err)
	}

	// LazyQuotes implementation to remove characters from https://stackoverflow.com/questions/31326659/golang-csv-error-bare-in-non-quoted-field
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
	reader.LazyQuotes = true

	for {

		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		for value := range record {
			fmt.Printf("%s\n", record[value])
			csvList = append(csvList, record[value])
		}
	}
	return csvList
}
