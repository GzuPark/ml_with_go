package main

import (
	"image/color"
	"log"
	"math"
	"os"
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var suffix = "sigmoid"

func main() {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Logistic Function"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "sigmoid(x)"

	sigmoidPlotter := plotter.NewFunction(func(x float64) float64 { return sigmoid(x) })
	sigmoidPlotter.Color = color.RGBA{B: 255, A: 255}

	p.Add(sigmoidPlotter)
	p.X.Min = -10
	p.X.Max = 10
	p.Y.Min = -0.1
	p.Y.Max = 1.1

	if err := p.Save(4*vg.Inch, 4*vg.Inch, plotPath("")); err != nil {
		log.Fatal(err)
	}
}

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func plotPath(name string) string {
	saveName := name + suffix + ".png"
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
			log.Fatal(err)
	}
	savePath := filepath.Join(dir, "result", saveName)

	return savePath
}
