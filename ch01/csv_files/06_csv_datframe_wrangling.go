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
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)
)

func main() {
	irisFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	filter := dataframe.F{
		Colname: "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	fmt.Println(versicolorDF)

	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})
	fmt.Println(versicolorDF)

	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})
	fmt.Println(versicolorDF)
}
