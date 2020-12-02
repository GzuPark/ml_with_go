package main

import (
	"encoding/csv"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var (
	fileName = "AirPassengers.csv"
	saveName = "diff_series.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "storage", "data", fileName)
	savePath = filepath.Join(os.Getenv("MLGO"), "storage", "data", saveName)
	suffix   = "diff_passengers_ts"
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	passengersVals := df.Col("AirPassengers").Float()
	timeVals := df.Col("time").Float()

	pts := make(plotter.XYs, df.Nrow()-1)

	var differenced [][]string
	differenced = append(differenced, []string{"time", "differenced_passengers"})

	for i := 1; i < len(passengersVals); i++ {
		pts[i-1].X = timeVals[i]
		pts[i-1].Y = passengersVals[i] - passengersVals[i-1]
		differenced = append(differenced, []string{
			strconv.FormatFloat(timeVals[i], 'f', -1, 64),
			strconv.FormatFloat(passengersVals[i]-passengersVals[i-1], 'f', -1, 64),
		})
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "time"
	p.Y.Label.Text = "differenced passengers"
	p.Add(plotter.NewGrid())

	l, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	p.Add(l)

	if err := p.Save(10*vg.Inch, 4*vg.Inch, plotPath("")); err != nil {
		log.Fatal(err)
	}

	f, err = os.Create(savePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(differenced)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
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
