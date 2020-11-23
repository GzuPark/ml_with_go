package main

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"github.com/go-gota/gota/dataframe"
)

func main() {
	f, err := os.Open("data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	sepalLength := df.Col("petal_length").Float()

	minVal := floats.Min(sepalLength)
	maxVal := floats.Max(sepalLength)

	rangeVal := maxVal - minVal

	varianceVal := stat.Variance(sepalLength, nil)
	stdDevVal := stat.StdDev(sepalLength, nil)

	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)

	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Min value: %.2f\n", minVal)
	fmt.Printf("Max value: %.2f\n", maxVal)
	fmt.Printf("Range value: %.2f\n", rangeVal)
	fmt.Printf("Variance value: %.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %.2f\n", stdDevVal)
	fmt.Printf("25 Quantile value: %.2f\n", quant25)
	fmt.Printf("50 Quantile value: %.2f\n", quant50)
	fmt.Printf("75 Quantile value: %.2f\n\n", quant75)
}
