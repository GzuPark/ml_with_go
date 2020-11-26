package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

var (
	testName = "advertising_test.csv"
	testPath = filepath.Join(os.Getenv("MLGO"), "data", testName)
)

const (
	intercept = 2.949
	coefTV = 0.047
	coefRadio = 0.180
	coefNewpaper = -0.001
)

func main() {
	f, err := os.Open(testPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var mAE float64

	for i, record := range testData {
		if i == 0 {
			continue
		}

		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		newspaperVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		yPredicted := predict(tvVal, radioVal, newspaperVal)
		mAE += math.Abs(yObserved - yPredicted) / float64(len(testData))
	}

	fmt.Printf("\nMAE = %.2f\n\n", mAE)
}

func predict(tv, radio, newspaper float64) float64 {
	return intercept + coefTV * tv + coefRadio * radio + coefNewpaper * newspaper
}
