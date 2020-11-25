package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	intercept = 0.00816
	coeff1    = 0.23495
	coeff2    = -0.17368
)

func main() {
	f, err := os.Open("../data/log_diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	passengers, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var predictions []float64

	for i, _ := range passengers {
		if i == 0 || i == 1 || i == 2 {
			continue
		}

		lagOne, err := strconv.ParseFloat(passengers[i - 1][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		lagTwo, err := strconv.ParseFloat(passengers[i - 2][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		predictions = append(predictions, intercept + coeff1 * lagOne + coeff2 * lagTwo)
	}

	f, err = os.Open("../data/AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader = csv.NewReader(f)
	reader.FieldsPerRecord = 2
	
	passengers, err = reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	ptsObs := make(plotter.XYs, len(predictions))
	ptsPred := make(plotter.XYs, len(predictions))

	var mAE    float64
	var cumSum float64

	for i := 4; i <= len(passengers) - 1; i++ {
		observed, err := strconv.ParseFloat(passengers[i][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		date, err := strconv.ParseFloat(passengers[i][0], 64)
		if err != nil {
			log.Fatal(err)
		}

		cumSum += predictions[i - 4]

		predicted := math.Exp(math.Log(observed) + cumSum)

		mAE += math.Abs(observed - predicted) / float64(len(predictions))

		ptsObs[i - 4].X  = date
		ptsPred[i - 4].X = date
		ptsObs[i - 4].Y  = observed
		ptsPred[i - 4].Y = predicted
	}

	fmt.Printf("\nMAE = %.2f\n\n", mAE)

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "time"
	p.Y.Label.Text = "passengers"
	p.Add(plotter.NewGrid())

	lObs, err := plotter.NewLine(ptsObs)
	if err != nil {
		log.Fatal(err)
	}
	lObs.LineStyle.Width = vg.Points(1)

	lPred, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	lPred.LineStyle.Width = vg.Points(1)
	lPred.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	p.Add(lObs, lPred)
	p.Legend.Add("Observed", lObs)
	p.Legend.Add("Predicted", lPred)

	if err := p.Save(10*vg.Inch, 4*vg.Inch, "result/predicted_passengers_ts.png"); err != nil {
		log.Fatal(err)
	}
}
