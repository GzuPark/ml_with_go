package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	fileName = "loan_data.csv"
	saveName = "clean_loan_data.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)
	savePath = filepath.Join(os.Getenv("MLGO"), "data", saveName)
)

const (
	scoreMax = 830.0
	scoreMin = 640.0
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	rawData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	f, err = os.Create(savePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	for idx, record := range rawData {
		if idx == 0 {
			if err := w.Write([]string{"FICO_score", "class"}); err != nil {
				log.Fatal(err)
			}
			continue
		}

		outRecord := make([]string, 2)

		score, err := strconv.ParseFloat(strings.Split(record[0], "-")[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		outRecord[0] = strconv.FormatFloat((score - scoreMin) / (scoreMax - scoreMin), 'f', 4, 64)

		rate, err := strconv.ParseFloat(strings.TrimSuffix(record[1], "%"), 64)
		if err != nil {
			log.Fatal(err)
		}

		if rate <= 12.0 {
			outRecord[1] = "1.0"

			if err := w.Write(outRecord); err != nil {
				log.Fatal(err)
			}
			continue
		}

		outRecord[1] = "0.0"

		if err := w.Write(outRecord); err != nil {
			log.Fatal(err)
		}
	}

	// buffer에 저장된 데이터를 표준 출력에 write
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
