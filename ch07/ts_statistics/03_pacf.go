package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
)

var (
	fileName = "AirPassengers.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "storage", "data", fileName)
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	passengers := df.Col("AirPassengers").Float()

	fmt.Println("Partial Autocorrelation:")
	for i := 1; i < 11; i++ {
		pac := pacf(passengers, i)
		fmt.Printf("Lag %d period: %.2f\n", i, pac)
	}
}

func pacf(x []float64, lag int) float64 {
	var r regression.Regression
	r.SetObserved("x")

	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	xAdj := x[lag:len(x)]

	for i, xVal := range xAdj {
		laggedVariables := make([]float64, lag)

		for idx := 1; idx <= lag; idx++ {
			laggedVariables[idx-1] = x[lag+i-idx]
		}

		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	r.Run()

	return r.Coeff(lag)
}
