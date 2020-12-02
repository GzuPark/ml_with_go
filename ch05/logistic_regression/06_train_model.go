package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gonum.org/v1/gonum/mat"
)

var (
	trainingName = "clean_loan_training.csv"
	trainingPath = filepath.Join(os.Getenv("MLGO"), "storage", "data", trainingName)
)

func main() {
	f, err := os.Open(trainingPath)
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

	featureData := make([]float64, 2*(len(rawData)-1))
	labels := make([]float64, len(rawData)-1)

	var featureIndex int

	for idx, record := range rawData {
		if idx == 0 {
			continue
		}

		featureVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		featureData[featureIndex] = featureVal
		featureData[featureIndex+1] = 1.0
		featureIndex += 2

		labelVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		labels[idx-1] = labelVal
	}

	features := mat.NewDense(len(rawData)-1, 2, featureData)
	weights := logisticRegression(features, labels, 1000, 0.3)

	formula := "p = 1 / ( 1 + exp(- m1 * FICO.score - m2) )"
	fmt.Printf("\n%s\n\nm1 = %.2f\nm2 = %.2f\n\n", formula, weights[0], weights[1])
}

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func logisticRegression(features *mat.Dense, labels []float64, numSteps int, learningRate float64) []float64 {
	_, numWeights := features.Dims()
	weights := make([]float64, numWeights)

	// reference
	s := rand.NewSource(time.Now().UnixNano())
	// reproducibility
	// s := rand.NewSource(42)
	r := rand.New(s)

	for idx, _ := range weights {
		weights[idx] = r.Float64()
	}

	for i := 0; i < numSteps; i++ {
		var sumError float64

		for idx, label := range labels {
			featureRow := mat.Row(nil, idx, features)

			pred := sigmoid(featureRow[0]*weights[0] + featureRow[1]*weights[1])
			predError := label - pred
			sumError += math.Pow(predError, 2)

			for j := 0; j < len(featureRow); j++ {
				weights[j] += learningRate * predError * pred * (1 - pred) * featureRow[j]
			}
		}
	}

	return weights
}
