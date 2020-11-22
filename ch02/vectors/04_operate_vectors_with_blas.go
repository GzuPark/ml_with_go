package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/blas/blas64"
)

func main() {
	vectorA := mat.NewVecDense(3, []float64{11.0, 5.2, -1.3})
	vectorB := mat.NewVecDense(3, []float64{-7.2, 4.2, 5.1})

	dotProduct := mat.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %.2f\n", dotProduct)

	vectorA.ScaleVec(1.5, vectorA)
	// *mat.VecDense
	// https://github.com/gonum/gonum/blob/de0df9a81247732df6a1f83101c4afc35852c5ca/mat/vector.go#L84
	// blas64.Vector
	// https://github.com/gonum/gonum/blob/b29604be86deb4e614fdcf6a1778f9f1cb7f071d/blas/blas64/blas64.go#L29
	fmt.Printf("Number of A : %v\n", vectorA.RawVector().N)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorA.RawVector().Data)
	fmt.Printf("Incremental of A: %v\n", vectorA.RawVector().Inc)

	normB := blas64.Nrm2(vectorB.RawVector())
	fmt.Printf("The norm/length of B is: %.2f\n", normB)
}
