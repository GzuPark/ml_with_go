package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/stat"
)

func main() {
	f, err := os.Open("../data/AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	passengers := df.Col("AirPassengers").Float()

	fmt.Println("Autocorrelation:")
	for i := 1; i < 11; i++ {
		ac := acf(passengers, i)
		fmt.Printf("Lag %d period: %.2f\n", i, ac)
	}
}

func acf(x []float64, lag int) float64 {
	xAdj := x[lag:len(x)]
	xLag := x[0:len(x) - lag]

	var numerator   float64
	var denominator float64

	xBar := stat.Mean(x, nil)

	for idx, xVal := range xAdj {
		numerator += ((xVal - xBar) * (xLag[idx] - xBar))
	}

	for _, xVal := range x {
		denominator += math.Pow(xVal - xBar, 2)
	}

	return numerator / denominator
}
