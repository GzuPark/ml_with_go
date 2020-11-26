package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
)

var (
	fileName = "log_diff_series.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)
	suffix = "log_diff_pacf"
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
	p.Title.Text = "Partial Autocorrelation for log(differenced AirPassengers)"
	p.X.Label.Text = "Lag"
	p.Y.Label.Text = "PACF"
	p.Y.Min = 15
	p.Y.Max = -1

	w := vg.Points(3)

	numLags := 20
	pts := make(plotter.Values, numLags)

	for i := 1; i < numLags; i++ {
		pts[i-1] = pacf(passengers, i)
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

func pacf(x []float64, lag int) float64 {
	var r regression.Regression
	r.SetObserved("x")

	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	xAdj := x[lag:len(x)]

	for i, xVal := range xAdj {
		laggedVariables := make([]float64, lag)

		for idx := 1; idx <= lag; idx++ {
			laggedVariables[idx - 1] = x[lag + i - idx]
		}

		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	r.Run()

	return r.Coeff(lag)
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
