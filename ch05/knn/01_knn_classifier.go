package main

import (
	"fmt"
	"log"
	"math"
	// go get github.com/sjwhitworth/golearn
	// cd $GOPATH/src/github.com/sjwhitworth/golearn
	// git reset --hard 6fed29ee
	// go get ./...
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	rawData, err := base.ParseCSVToInstances("../data/iris.csv", true)
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
