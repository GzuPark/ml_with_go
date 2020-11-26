package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

var (
	testName = "clean_loan_test.csv"
	testPath = filepath.Join(os.Getenv("MLGO"), "data", testName)
)

const (
	m1 = 13.65
	m2 = -4.89
)

func main() {
	f, err := os.Open(testPath)
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
			line++
			continue
		}

		observedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		score, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal := predict(score)

		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	var truePosNeg int

	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}

	accuracy := float64(truePosNeg) / float64(len(observed))
	fmt.Printf("\nAccuracy = %.2f\n\n", accuracy)
}

func predict(score float64) float64 {
	p := 1 / (1 + math.Exp(-1 * m1 * score -1 * m2))

	if p >= 0.5 {
		return 1.0
	}

	return 0.0
}
