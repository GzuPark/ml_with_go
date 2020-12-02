package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
}

var (
	fileName = "iris_mixed_types.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "storage", "data", fileName)
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	var csvData []CSVRecord

	line := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		var csvRecord CSVRecord

		for idx, value := range record {
			if idx == 4 {
				if value == "" {
					log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
					csvRecord.ParseError = fmt.Errorf("Empty string value")
					break
				}

				csvRecord.Species = value
				continue
			}

			var floatValue float64

			if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
				csvRecord.ParseError = fmt.Errorf("Could not parse float")
				break
			}

			switch idx {
			case 0:
				csvRecord.SepalLength = floatValue
			case 1:
				csvRecord.SepalWidth = floatValue
			case 2:
				csvRecord.PetalLength = floatValue
			case 3:
				csvRecord.PetalWidth = floatValue
			}
		}

		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}

		line++
	}

	fmt.Printf("successfully parsed %d lines\n", len(csvData))
}
