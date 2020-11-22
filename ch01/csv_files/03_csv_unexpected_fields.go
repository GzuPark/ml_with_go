package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("data/iris_unexpected_fields.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	// row 마다 읽을 column(=field) 5개로 설정
	// useful at validating structured data format
	reader.FieldsPerRecord = 5

	var rawCSVData [][]string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println(err)
			continue
		}

		rawCSVData = append(rawCSVData, record)
	}

	fmt.Printf("parsed %d lines successfully\n", len(rawCSVData))
}
