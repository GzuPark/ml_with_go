package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	data := []float64{1.2, -5.7, -2.4, 7.3}

	a := mat.NewDense(2, 2, data)

	fa := mat.Formatted(a, mat.Prefix(" "))
	fmt.Printf("A = %v\n\n", fa)
}
