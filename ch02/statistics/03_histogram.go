package main

import (
	"fmt"
	"log"
	"os"
	// go get gonum.org/v1/plot/...
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
)

func main() {
	f, err := os.Open("data/iris.csv")
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

			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				log.Fatal(err)
			}
		}
	}
}
