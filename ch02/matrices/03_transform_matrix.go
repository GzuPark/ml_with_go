package main

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	ft := mat.Formatted(a.T(), mat.Prefix(" "))
	fmt.Printf("a^T = %v\n\n", ft)

	det_a := mat.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", det_a)

	// https://github.com/gonum/gonum/blob/b29604be86deb4e614fdcf6a1778f9f1cb7f071d/mat/dense_arithmetic.go#L213
	// (example) https://godoc.org/gonum.org/v1/gonum/mat#ex-Dense-Inverse
	var aInverse mat.Dense
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}

	fi := mat.Formatted(&aInverse, mat.Prefix(" "), mat.Squeeze())
	fmt.Printf("a^-1 = %v\n\n", fi)
}
