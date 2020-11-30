package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sajari/regression"
)

type ModelInfo struct {
	Intercept    float64           `json:"intercept"`
	Coefficients []CoefficientInfo `json:"coefficients"`
}

type CoefficientInfo struct {
	Name        string  `json:"name"`
	Coefficient float64 `json:"coefficient"`
}

func main() {
	inDirPtr := flag.String("inDir", "", "The directory containing the training data")
	outDirPtr := flag.String("outDir", "", "The output directory")
	flag.Parse()

	f, err := os.Open(filepath.Join(*inDirPtr, "diabetes_training.csv"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 11
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var r regression.Regression
	r.SetObserved("diabetes progression")
	r.SetVar(0, "bmi")

	for i, record := range trainingData {
		if i == 0 {
			continue
		}

		yVal, err := strconv.ParseFloat(record[10], 64)
		if err != nil {
			log.Fatal(err)
		}

		bmiVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		r.Train(regression.DataPoint(yVal, []float64{bmiVal}))
	}

	r.Run()

	fmt.Printf("\nRegression Formula: \n%v\n\n", r.Formula)

	modelInfo := ModelInfo{
		Intercept: r.Coeff(0),
		Coefficients: []CoefficientInfo{
			{
				Name:        "bmi",
				Coefficient: r.Coeff(1),
			},
		},
	}

	outputData, err := json.MarshalIndent(modelInfo, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(filepath.Join(*outDirPtr, "model.json"), outputData, 0644); err != nil {
		log.Fatal(err)
	}
}
