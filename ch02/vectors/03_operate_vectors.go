package main

import (
	"fmt"

	"gonum.org/v1/gonum/floats"
)

func main() {
	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}

	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %.2f\n", dotProduct)

	floats.Scale(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorA)

	normB := floats.Norm(vectorB, 2)
	fmt.Printf("The norm/length of B is: %.2f\n", normB)
}
