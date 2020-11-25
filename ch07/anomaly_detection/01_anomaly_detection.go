package main

import (
	"fmt"
	"log"

	"github.com/lytics/anomalyzer"
)

const (
	exampleOne   = 15.2
	exampleTwo   = 0.43
)

func main() {
	conf := &anomalyzer.AnomalyzerConf{
		Sensitivity: 0.1,
		UpperBound:  5,
		LowerBound:  anomalyzer.NA, // ignore lower bound
		ActiveSize:  1,
		NSeasons:    4,
		Methods:     []string{"diff", "fence", "highrank", "lowrank", "magnitude"},
	}

	ts := []float64{0.1, 0.2, 0.5, 0.12, 0.38, 0.9, 0.74}

	model, err := anomalyzer.NewAnomalyzer(conf, ts)
	if err != nil {
		log.Fatal(err)
	}

	example := []float64{exampleOne, exampleTwo}

	for _, ex := range example {
		prob := model.Push(ex)
		fmt.Printf("Probability of %f being anomalous: %.2f\n", ex, prob)	
	}
}
