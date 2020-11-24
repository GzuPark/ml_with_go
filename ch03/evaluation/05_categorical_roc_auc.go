package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/integrate"
)

func main() {
	scores := []float64{0.1, 0.35, 0.4, 0.8}
	classes := []bool{true, false, true, false}

	// (example) https://godoc.org/gonum.org/v1/gonum/stat#ex-ROC--AUC
	// tpr: true positive ratio
	// fpr: false positive ratio
	tpr, fpr, _ := stat.ROC(nil, scores, classes, nil)

	auc := integrate.Trapezoidal(fpr, tpr)

	fmt.Printf("\ntrue positive rate: %v\n", tpr)
	fmt.Printf("false positive rate: %v\n", fpr)
	fmt.Printf("auc: %v\n\n", auc)
}
