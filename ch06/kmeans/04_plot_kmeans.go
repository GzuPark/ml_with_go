package main

import (
	"image/color"
	"log"
	"os"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
)

var (
	cOne = []float64{180.02, 18.29}
	cTwo = []float64{50.05, 8.83}
)

func main() {
	f, err := os.Open("../data/fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	yVals := df.Col("Distance_Feature").Float()

	var clusterOne [][]float64
	var clusterTwo [][]float64

	for i, xVal := range df.Col("Speeding_Feature").Float() {
		distanceOne := floats.Distance([]float64{yVals[i], xVal}, cOne, 2)
		distanceTwo := floats.Distance([]float64{yVals[i], xVal}, cTwo, 2)
		if distanceOne < distanceTwo {
			clusterOne = append(clusterOne, []float64{xVal, yVals[i]})
			continue
		}
		clusterTwo = append(clusterTwo, []float64{xVal, yVals[i]})
	}

	ptsOne := make(plotter.XYs, len(clusterOne))
	ptsTwo := make(plotter.XYs, len(clusterTwo))

	for i, point := range clusterOne {
		ptsOne[i].X = point[0]
		ptsOne[i].Y = point[1]
	}

	for i, point := range clusterTwo {
		ptsTwo[i].X = point[0]
		ptsTwo[i].Y = point[1]
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "Speeding"
	p.Y.Label.Text = "Distance"
	p.Add(plotter.NewGrid())

	sOne, err := plotter.NewScatter(ptsOne)
	if err != nil {
		log.Fatal(err)
	}
	sOne.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	sOne.GlyphStyle.Radius = vg.Points(3)

	sTwo, err := plotter.NewScatter(ptsTwo)
	if err != nil {
		log.Fatal(err)
	}
	sTwo.GlyphStyle.Color = color.RGBA{B: 255, A: 255}
	sTwo.GlyphStyle.Radius = vg.Points(3)

	p.Add(sOne, sTwo)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "result/fleet_data_clusters.png"); err != nil {
		log.Fatal(err)
	}
}
