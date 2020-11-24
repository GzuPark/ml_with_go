package main

import (
	"fmt"
	"log"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/filters"
	"github.com/sjwhitworth/golearn/naive"
)

func main() {
	trainingData, err := base.ParseCSVToInstances("../data/training.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	cls := naive.NewBernoulliNBClassifier()
	cls.Fit(convertToBinary(trainingData))

	testData, err := base.ParseCSVToInstances("../data/test.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	predictions, err := cls.Predict(convertToBinary(testData))
	if err != nil {
		log.Fatal(err)
	}

	cm, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		log.Fatal(err)
	}

	accuracy := evaluation.GetAccuracy(cm)
	fmt.Printf("\nAccuracy : %.2f\n\n", accuracy)
}

// label을 binary로 변환
func convertToBinary(src base.FixedDataGrid) base.FixedDataGrid {
	b := filters.NewBinaryConvertFilter()
	attrs := base.NonClassAttributes(src)
	for _, a := range attrs {
		b.AddAttribute(a)
	}
	b.Train()
	ret := base.NewLazilyFilteredInstances(src, b)
	return ret
}
