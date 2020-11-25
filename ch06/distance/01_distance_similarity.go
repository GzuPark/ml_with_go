package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func main() {
	distance := floats.Distance([]float64{1, 2}, []float64{3, 4}, 2)
	fmt.Printf("\nDistance: %.2f\n\n", distance)
}
