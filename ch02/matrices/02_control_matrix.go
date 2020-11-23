package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	data := []float64{1.2, -5.7, -2.4, 7.3}

	a := mat.NewDense(2, 2, data)

	val := a.At(0, 1)
	fmt.Printf("The value of a at (0, 1) is: %.2f\n\n", val)

	col := mat.Col(nil, 0, a)
	fmt.Printf("The values in the 1st column are: %v\n\n", col)

	row := mat.Row(nil, 1, a)
	fmt.Printf("The values in the 2nd row are: %v\n\n", row)

	a.Set(1, 1, 11.2)
	a.SetRow(0, []float64{14.3, -4.2})
	a.SetCol(0, []float64{1.7, -0.3})

	fa := mat.Formatted(a, mat.Prefix(" "))
	fmt.Printf("A = %v\n\n", fa)
}
