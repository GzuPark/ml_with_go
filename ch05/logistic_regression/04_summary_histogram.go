package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
)

var (
	fileName = "clean_loan_data.csv"
	filePath = filepath.Join(os.Getenv("MLGO"), "data", fileName)
	suffix = "_hist"
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)
	
	summary := df.Describe()
	fmt.Println(summary)

	for _, colName := range df.Names() {
		plotVals := make(plotter.Values, df.Nrow())
		for i, floatVal := range df.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		p, err := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of %s", colName)

		h, err := plotter.NewHist(plotVals, 16)
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

func plotPath(name string) string {
	saveName := name + suffix + ".png"
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
			log.Fatal(err)
	}
	savePath := filepath.Join(dir, "result", saveName)

	return savePath
}
