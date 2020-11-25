package main

import (
	"log"
	"math"
	"os"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
)

func main() {
	f, err := os.Open("../data/AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	passengers := df.Col("AirPassengers").Float()

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Autocorrelation for AirPassengers"
	p.X.Label.Text = "Lag"
	p.Y.Label.Text = "ACF"
	p.Y.Min = 0
	p.Y.Max = 1

	w := vg.Points(3)

	numLags := 20
	pts := make(plotter.Values, numLags)

	for i := 1; i < numLags; i++ {
		pts[i-1] = acf(passengers, i)
	}

	b, err := plotter.NewBarChart(pts, w)
	if err != nil {
		log.Fatal(err)
	}
	b.LineStyle.Width = vg.Length(0)
	b.Color = plotutil.Color(1)

	p.Add(b)

	if err := p.Save(8*vg.Inch, 4*vg.Inch, "result/acf.png"); err != nil {
		log.Fatal(err)
	}
}

func acf(x []float64, lag int) float64 {
	xAdj := x[lag:len(x)]
	xLag := x[0:len(x) - lag]

	var numerator   float64
	var denominator float64

	xBar := stat.Mean(x, nil)

	for idx, xVal := range xAdj {
		numerator += ((xVal - xBar) * (xLag[idx] - xBar))
	}

	for _, xVal := range x {
		denominator += math.Pow(xVal - xBar, 2)
	}

	return numerator / denominator
}
