package main

import (
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
	suffix   = "boxplots"
)

func main() {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Box plots"
	p.Y.Label.Text = "Values"

	w := vg.Points(50)

	for idx, colName := range df.Names() {
		if colName != "species" {
			v := make(plotter.Values, df.Nrow())
			for i, floatVal := range df.Col(colName).Float() {
				v[i] = floatVal
			}

			b, err := plotter.NewBoxPlot(w, float64(idx), v)
			if err != nil {
				log.Fatal(err)
			}
			p.Add(b)
		}
	}

	// 지정된 이름을 사용해 X-axis 이름 설정
	p.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")

	if err := p.Save(6*vg.Inch, 8*vg.Inch, plotPath("")); err != nil {
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
