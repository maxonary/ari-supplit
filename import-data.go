package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var inputFile = "./material/sneaker-numbers.csv"

func main() {
	fmt.Println("hello world")

	csvFile, err := os.Open(inputFile)

	if err != nil {

		log.Fatal(err)
	}

	// LazyQuotes implementation from https://stackoverflow.com/questions/31326659/golang-csv-error-bare-in-non-quoted-field
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
		}
	}
}
