package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

var (
	trainingName = "advertising_training.csv"
	trainingPath = filepath.Join(os.Getenv("MLGO"), "storage", "data", trainingName)
)

func main() {
	f, err := os.Open(trainingPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4

	rawData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	featureData := make([]float64, 4*len(rawData))
	yData := make([]float64, len(rawData))

	var featureIndex int
	var yIndex int

	for idx, record := range rawData {
		if idx == 0 {
			continue
		}

		for i, val := range record {
			valParsed, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal("Could not parse float value (index %d)", i)
			}

			if i < 3 {
				if i == 0 {
					featureData[featureIndex] = 1
					featureIndex++
				}

				featureData[featureIndex] = valParsed
				featureIndex++
			}

			if i == 3 {
				yData[yIndex] = valParsed
				yIndex++
			}
		}
	}

	features := mat.NewDense(len(rawData), 4, featureData)
	y := mat.NewVecDense(len(rawData), yData)

	if features != nil && y != nil {
		fmt.Println("Matrices formed for ridge regression")
	}
}
