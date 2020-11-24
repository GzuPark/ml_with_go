package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Sigmoid(1.0) = %.3f\n", sigmoid(1.0))
}

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
