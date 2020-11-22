package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("myfile.csv")
	if err != nil {
		log.Fatal(err)
	}
	
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var intMax int
	for _, record := range records {
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		if intVal > intMax {
			intMax = intVal
		}
	}

	fmt.Println(intMax)
}
