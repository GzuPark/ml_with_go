package main

import (
	"image/color"
	"log"
	"os"
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
)

var (
	fileName = "advertising.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)
	suffix = "_scatter"
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	yVals := df.Col("Sales").Float()

	for _, colName := range df.Names() {
		pts := make(plotter.XYs, df.Nrow())

		for i, floatVal := range df.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}

		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
		s.GlyphStyle.Radius = vg.Points(3)

		p.Add(s)

		if err := p.Save(4*vg.Inch, 4*vg.Inch, plotPath(colName)); err != nil {
			log.Fatal(err)
		}
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
