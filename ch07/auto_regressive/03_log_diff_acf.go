package main

import (
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var (
	fileName = "log_diff_series.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "storage", "data", fileName)
	suffix   = "log_diff_acf"
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	passengers := df.Col("log_differenced_passengers").Float()

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Autocorrelation for log(differenced AirPassengers)"
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

	if err := p.Save(8*vg.Inch, 4*vg.Inch, plotPath("")); err != nil {
		log.Fatal(err)
	}
}

func acf(x []float64, lag int) float64 {
	xAdj := x[lag:len(x)]
	xLag := x[0 : len(x)-lag]

	var numerator float64
	var denominator float64

	xBar := stat.Mean(x, nil)

	for idx, xVal := range xAdj {
		numerator += ((xVal - xBar) * (xLag[idx] - xBar))
	}

	for _, xVal := range x {
		denominator += math.Pow(xVal-xBar, 2)
	}

	return numerator / denominator
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
