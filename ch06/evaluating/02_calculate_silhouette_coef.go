package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gonum.org/v1/gonum/floats"
	"github.com/go-gota/gota/dataframe"
)

var (
	fileName = "iris.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)
)

type centroid []float64

func main() {
	f, err := os.Open(filePath)
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
	clusters := make(map[string]dataframe.DataFrame)

	for _, species := range speciesNames {
		filter := dataframe.F{
			Colname: "species",
			Comparator: "==",
			Comparando: species,
		}
		filtered := df.Filter(filter)
		clusters[species] = filtered

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

	labels := df.Col("species").Records()
	floatColumns := []string{
		"sepal_length",
		"sepal_width",
		"petal_length",
		"petal_width",
	}

	var silhouette float64

	for idx, label := range labels {
		var a float64

		for i := 0; i < clusters[label].Nrow(); i++ {
			current := dfFloatRow(df, floatColumns, idx)
			other := dfFloatRow(clusters[label], floatColumns, i)

			a += floats.Distance(current, other, 2) / float64(clusters[label].Nrow())
		}

		var otherCluster string
		var distanceToCluster float64

		for _, species := range speciesNames {
			if species == label {
				continue
			}

			distanceForThisCluster := floats.Distance(centroids[label], centroids[species], 2)

			if distanceToCluster == 0.0 || distanceForThisCluster < distanceToCluster {
				otherCluster = species
				distanceToCluster = distanceForThisCluster
			}
		}

		var b float64

		for i := 0; i < clusters[otherCluster].Nrow(); i++ {
			current := dfFloatRow(df, floatColumns, idx)
			other := dfFloatRow(clusters[otherCluster], floatColumns, i)

			b += floats.Distance(current, other, 2) / float64(clusters[otherCluster].Nrow())
		}

		if a > b {
			silhouette += ((b - a) / a) / float64(len(labels))
		}
		silhouette += ((b - a) / b) / float64(len(labels))
	}

	fmt.Printf("\nAverage Silhouette Coefficient: %.2f\n\n", silhouette)
}

func dfFloatRow(df dataframe.DataFrame, names []string, idx int) []float64 {
	var row []float64
	for _, name := range names {
		row = append(row, df.Col(name).Float()[idx])
	}

	return row
}
