package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/mash/gokmeans"
)

func main() {
	f, err := os.Open("../data/fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 3

	var data []gokmeans.Node

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if record[0] == "Driver_ID" {
			continue
		}

		var point []float64

		for i := 1; i < 3; i++ {
			val, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				log.Fatal(err)
			}

			point = append(point, val)
		}

		data = append(data, gokmeans.Node{point[0], point[1]})
	}

	success, centroids := gokmeans.Train(data, 2, 50)
	if !success {
		log.Fatal("Could not generate clusters")
	}

	fmt.Println("The centroids for our clusters are:")
	for _, centroid := range centroids {
		fmt.Println(centroid)
	}
}
