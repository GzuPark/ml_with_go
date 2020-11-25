package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

type centroid []float64

func main() {
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	speciesNames := []string{
		"Iris-setosa",
		"Iris-versicolor",
		"Iris-virginica",
	}

	centroids := make(map[string]centroid)

	for _, species := range speciesNames {
		filter := dataframe.F{
			Colname: "species",
			Comparator: "==",
			Comparando: species,
		}
		filtered := df.Filter(filter)

		summary := filtered.Describe()

		var c centroid
		for _, feature := range df.Names() {
			if feature == "column" || feature == "species" {
				continue
			}
			c = append(c, summary.Col(feature).Float()[0])
		}

		centroids[species] = c
	}

	for _, species := range speciesNames {
		fmt.Printf("%s centroid: %v\n", species, centroids[species])
	}
}
