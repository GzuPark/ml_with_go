package main

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	b := mat.NewDense(3, 3, []float64{8, 9, 10, 1, 4, 2, 9, 0, 2})
	c := mat.NewDense(3, 2, []float64{3, 2, 1, 4, 0, 8})

	var d mat.Dense
	d.Add(a, b)
	fd := mat.Formatted(&d, mat.Prefix(" "))
	fmt.Printf("d = a + b =\n %v\n\n", fd)

	var f mat.Dense
	f.Mul(a, c)
	ff := mat.Formatted(&f, mat.Prefix(" "))
	fmt.Printf("f = a . c =\n %v\n\n", ff)

	var g mat.Dense
	g.Pow(a, 5)
	fg := mat.Formatted(&g, mat.Prefix(" "))
	fmt.Printf("g = a^5 =\n %v\n\n", fg)

	// matrix cell마다 함수 적용
	var h mat.Dense
	sqrt := func(_, _ int, v float64) float64 { return math.Sqrt(v) }
	h.Apply(sqrt, a)
	fh := mat.Formatted(&h, mat.Prefix(" "))
	fmt.Printf("h = sqrt(a) =\n %0.4v\n\n", fh)
}
