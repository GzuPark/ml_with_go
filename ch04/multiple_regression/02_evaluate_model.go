package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sajari/regression"
)

var (
	trainingName = "advertising_training.csv"
	testName     = "advertising_test.csv"
	trainingPath = filepath.Join(os.Getenv("MLGO"), "storage", "data", trainingName)
	testPath     = filepath.Join(os.Getenv("MLGO"), "storage", "data", testName)
)

func main() {
	f, err := os.Open(trainingPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")
	r.SetVar(1, "Radio")

	for i, record := range trainingData {
		if i == 0 {
			continue
		}

		yVal, err := strconv.ParseFloat(record[3], 64)
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

		r.Train(regression.DataPoint(yVal, []float64{tvVal, radioVal}))
	}

	r.Run()
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)

	f, err = os.Open(testPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader = csv.NewReader(f)

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

		yPredicted, err := r.Predict([]float64{tvVal, radioVal})
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}

	fmt.Printf("MAE = %.2f\n\n", mAE)
}
