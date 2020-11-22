package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	irisFile, err := os.Open("data/iris_labeled.csv")
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
