package main

import (
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
)

func main() {
	f, err := os.Open("../data/fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	yVals := df.Col("Distance_Feature").Float()

	pts := make(plotter.XYs, df.Nrow())

	for i, floatVal := range df.Col("Speeding_Feature").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "Speeding"
	p.Y.Label.Text = "Distance"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	p.Add(s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "result/fleet_data_scatter.png"); err != nil {
		log.Fatal(err)
	}
}
