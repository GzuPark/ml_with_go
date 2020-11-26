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
	fileName = "fleet_data.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)

	cOne = []float64{180.02, 18.29}
	cTwo = []float64{50.05, 8.83}
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	distances := df.Col("Distance_Feature").Float()

	var clusterOne [][]float64
	var clusterTwo [][]float64

	for i, speed := range df.Col("Speeding_Feature").Float() {
		distanceOne := floats.Distance([]float64{distances[i], speed}, cOne, 2)
		distanceTwo := floats.Distance([]float64{distances[i], speed}, cTwo, 2)
		if distanceOne < distanceTwo {
			clusterOne = append(clusterOne, []float64{distances[i], speed})
			continue
		}
		clusterTwo = append(clusterTwo, []float64{distances[i], speed})
	}

	fmt.Printf("\nCluster 1 Metric: %.2f\n", withinClusterMean(clusterOne, cOne))
	fmt.Printf("Cluster 2 Metric: %.2f\n\n", withinClusterMean(clusterTwo, cTwo))
}

func withinClusterMean(cluster [][]float64, centroid []float64) float64 {
	var meanDistance float64

	for _, point := range cluster {
		meanDistance += floats.Distance(point, centroid, 2) / float64(len(cluster))
	}

	return meanDistance
}
