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
	fileName = "log_diff_series.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	passengers := df.Col("log_differenced_passengers").Float()

	coeffs, intercept := autoregressive(passengers, 2)

	// lag1: 1주기 이전
	// lag2: 2주기 이전
	fmt.Printf("\nlog(x(t)) - log(x(t-1)) = %.5f + lag1*%.5f + lag2*%.5f\n\n", intercept, coeffs[0], coeffs[1])
}

func autoregressive(x []float64, lag int) ([]float64, float64) {
	var r regression.Regression
	r.SetObserved("x")

	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	xAdj := x[lag:len(x)]

	for i, xVal := range xAdj {
		laggedVariables := make([]float64, lag)

		for idx := 1; idx <= lag; idx++ {
			laggedVariables[idx - 1] = x[lag + i - idx]
		}

		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	r.Run()

	var coeff []float64

	for i := 1; i <= lag; i++ {
		coeff = append(coeff, r.Coeff(i))
	}

	return coeff, r.Coeff(0)
}
