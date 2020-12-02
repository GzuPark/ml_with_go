package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var (
	fileName       = "log_diff_series.csv"
	originFileName = "AirPassengers.csv"
	filePath       = filepath.Join(os.Getenv("MLGO"), "storage", "data", fileName)
	originFilePath = filepath.Join(os.Getenv("MLGO"), "storage", "data", originFileName)
	suffix         = "predicted_passengers_ts"
)

const (
	intercept = 0.00816
	coeff1    = 0.23495
	coeff2    = -0.17368
)

func main() {
	f, err := os.Open(filePath)
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

		lagOne, err := strconv.ParseFloat(passengers[i-1][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		lagTwo, err := strconv.ParseFloat(passengers[i-2][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		predictions = append(predictions, intercept+coeff1*lagOne+coeff2*lagTwo)
	}

	f, err = os.Open(originFilePath)
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

	var mAE float64
	var cumSum float64

	for i := 4; i <= len(passengers)-1; i++ {
		observed, err := strconv.ParseFloat(passengers[i][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		date, err := strconv.ParseFloat(passengers[i][0], 64)
		if err != nil {
			log.Fatal(err)
		}

		cumSum += predictions[i-4]

		predicted := math.Exp(math.Log(observed) + cumSum)

		mAE += math.Abs(observed-predicted) / float64(len(predictions))

		ptsObs[i-4].X = date
		ptsPred[i-4].X = date
		ptsObs[i-4].Y = observed
		ptsPred[i-4].Y = predicted
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

	if err := p.Save(10*vg.Inch, 4*vg.Inch, plotPath("")); err != nil {
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
