package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-gota/gota/dataframe"
)

var (
	fileName = "AirPassengers.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	fmt.Println(df)
}
