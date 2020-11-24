package main

import (
	"bufio"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	f, err := os.Open("../data/advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	trainingNum := (4 * df.Nrow()) / 5
	testNum := df.Nrow() / 5

	if trainingNum + testNum < df.Nrow() {
		trainingNum++
	}

	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	for i:= 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	for i:= 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	trainingDF := df.Subset(trainingIdx)
	testDF := df.Subset(testIdx)

	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	for idx, setName := range []string{"../data/training.csv", "../data/test.csv"} {
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		w := bufio.NewWriter(f)

		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
