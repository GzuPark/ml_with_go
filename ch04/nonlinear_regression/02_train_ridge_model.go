package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pa-m/sklearn/linear_model"
	"gonum.org/v1/gonum/mat"
)

func main() {
	f, err := os.Open("../data/training.csv")
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

	// (example) https://godoc.org/github.com/pa-m/sklearn/linear_model#ex-Ridge
	r := linearmodel.NewRidge()
	r.Tol = 1e-15
	r.Normalize = true
	r.Alpha = 0
	r.L1Ratio = 0.
	r.Fit(features, y)

	c1 := r.Coef.At(0, 0)
	c2 := r.Coef.At(1, 0)
	c3 := r.Coef.At(2, 0)
	c4 := r.Coef.At(3, 0)
	fmt.Printf("\nRegression formula:\n")
	fmt.Printf("y = %.3f + %.3f TV + %.3f Radio + %.3f Newspaper\n\n", c1, c2, c3, c4)
}
