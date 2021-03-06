package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var (
	fileName = "iris.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "storage", "data", fileName)
	suffix   = "_hist"
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	for _, colName := range df.Names() {
		if colName != "species" {
			v := make(plotter.Values, df.Nrow())
			for i, floatVal := range df.Col(colName).Float() {
				v[i] = floatVal
			}

			p, err := plot.New()
			if err != nil {
				log.Fatal(err)
			}
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			h.Normalize(1)

			p.Add(h)

			if err := p.Save(4*vg.Inch, 4*vg.Inch, plotPath(colName)); err != nil {
				log.Fatal(err)
			}
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
