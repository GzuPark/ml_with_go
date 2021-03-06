package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
)

var (
	fileName = "iris.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "storage", "data", fileName)
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	sepalLength := df.Col("petal_length").Float()

	meanVal := stat.Mean(sepalLength, nil)
	modeVal, modeCount := stat.Mode(sepalLength, nil)

	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %.2f\n", meanVal)
	fmt.Printf("Mode value: %.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
	fmt.Printf("Median value: %.2f\n\n", medianVal)
}
