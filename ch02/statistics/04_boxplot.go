package main

import (
	"log"
	"os"

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

	if err := p.Save(6*vg.Inch, 8*vg.Inch, "boxplots.png"); err != nil {
		log.Fatal(err)
	}
}
