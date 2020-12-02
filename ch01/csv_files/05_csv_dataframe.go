package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-gota/gota/dataframe"
)

var (
	fileName = "iris_labeled.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "storage", "data", fileName)
)

func main() {
	irisFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	fmt.Println(irisDF)
}
