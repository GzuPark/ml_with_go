package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

var (
	fileName = "iris.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)
)

func main() {
	rawData, err := base.ParseCSVToInstances(filePath, true)
	if err != nil {
		log.Fatal(err)
	}

	cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(rawData, cls, 5)
	if err != nil {
		log.Fatal(err)
	}

	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	fmt.Printf("\nAccuracy :\n %.2f (+/- %.2f)\n\n", mean, stdev*2)
}
