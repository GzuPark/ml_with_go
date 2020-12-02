package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var (
	fileName = "labeled.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "storage", "data", fileName)
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	var observed []int
	var predicted []int

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

		observedVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	classes := []int{0, 1, 2}

	for _, class := range classes {
		var truePos int
		var falsePos int
		var falseNeg int

		for idx, oVal := range observed {
			switch oVal {
			case class:
				if predicted[idx] == class {
					truePos++
					continue
				}
				falseNeg++
			default:
				if predicted[idx] == class {
					falsePos++
				}
			}
		}

		precision := float64(truePos) / float64(truePos+falsePos)
		recall := float64(truePos) / float64(truePos+falseNeg)

		fmt.Printf("\nPrecision (class %d) = %.2f\n", class, precision)
		fmt.Printf("Recall    (class %d) = %.2f\n", class, recall)
	}

	fmt.Println()
}
