package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/trees"
)

func main() {
	rawData, err := base.ParseCSVToInstances("../data/iris.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(42)

	cls := trees.NewID3DecisionTree(0.6)

	cv, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(rawData, cls, 5)
	if err != nil {
		log.Fatal(err)
	}

	mean, variance := evaluation.GetCrossValidatedMetric(cv, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	fmt.Printf("\nAccuracy :\n %.2f (+/- %.2f)\n\n", mean, stdev*2)
}
