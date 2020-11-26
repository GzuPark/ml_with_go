package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	fileName = "iris_without_header.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	// rows를 성공적으로 읽으면 데이터에 저장 
	var rawCSVData [][]string

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		rawCSVData = append(rawCSVData, record)
	}

	fmt.Println(rawCSVData)
}
