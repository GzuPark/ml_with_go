package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sajari/regression"
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
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var r regression.Regression
	// Y-axis
	r.SetObserved("Sales")
	// X-axis
	r.SetVar(0, "TV")

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

		r.Train(regression.DataPoint(yVal, []float64{tvVal}))
	}

	r.Run()
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)
}
