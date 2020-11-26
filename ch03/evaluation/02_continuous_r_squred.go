package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"gonum.org/v1/gonum/stat"
)

var (
	fileName = "continuous_data.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	var observed []float64
	var predicted []float64

	line := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if line == 1 {
			line ++
			continue
		}

		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	rSquared := stat.RSquaredFrom(observed, predicted, nil)

	fmt.Printf("\nR^2 = %.2f\n\n", rSquared)
}
