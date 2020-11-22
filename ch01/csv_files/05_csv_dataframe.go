package main

import (
	"fmt"
	"log"
	"os"
	// go get github.com/go-gota/gota/...
	"github.com/go-gota/gota/dataframe"
)

func main() {
	irisFile, err := os.Open("data/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	fmt.Println(irisDF)
}
