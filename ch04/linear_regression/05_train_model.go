package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	// go get github.com/sajari/regression@d629f2e9549b2bb5f7c5b2bcfa716f074feb2d82
	"github.com/sajari/regression"
)

func main() {
	f, err := os.Open("../data/training.csv")
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
		if err != nil{
			log.Fatal(err)
		}

		r.Train(regression.DataPoint(yVal, []float64{tvVal}))
	}

	r.Run()
	fmt.Printf("\nRegression Formula:\n%v\n\n", r.Formula)
}
