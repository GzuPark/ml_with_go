package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type ModelInfo struct {
	Intercept    float64           `json:"intercept"`
	Coefficients []CoefficientInfo `json:"coefficients"`
}

type CoefficientInfo struct {
	Name        string  `json:"name"`
	Coefficient float64 `json:"coefficient"`
}

type PredictionData struct {
	Prediction      float64           `json:"predicted_diabetes_progression"`
	IndependentVars []IndependentVars `json:"independent_variables"`
}

type IndependentVars struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func main() {
	inModelDirPtr := flag.String("inModelDir", "", "The directory containing the model")
	inVarDirPtr := flag.String("inVarDir", "", "The directory containing the input attributes")
	outDirPtr := flag.String("outDir", "", "The output directory")
	flag.Parse()

	f, err := ioutil.ReadFile(filepath.Join(*inModelDirPtr, "model.json"))
	if err != nil {
		log.Fatal(err)
	}

	var modelInfo ModelInfo
	if err := json.Unmarshal(f, &modelInfo); err != nil {
		log.Fatal(err)
	}

	if err := filepath.Walk(*inVarDirPtr, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		f, err := ioutil.ReadFile(filepath.Join(*inVarDirPtr, info.Name()))
		if err != nil {
			return err
		}

		var predictionData PredictionData
		if err := json.Unmarshal(f, &predictionData); err != nil {
			return err
		}

		if err := Predict(&modelInfo, &predictionData); err != nil {
			return err
		}

		outputData, err := json.MarshalIndent(predictionData, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		fileName := "predicted_" + info.Name()
		if err := ioutil.WriteFile(filepath.Join(*outDirPtr, fileName), outputData, 0644); err != nil {
			log.Fatal(err)
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func Predict(modelInfo *ModelInfo, predictionData *PredictionData) error {
	prediction := modelInfo.Intercept

	coeffs := make(map[string]float64)
	varNames := make([]string, len(modelInfo.Coefficients))
	for idx, coeff := range modelInfo.Coefficients {
		coeffs[coeff.Name] = coeff.Coefficient
		varNames[idx] = coeff.Name
	}

	varVals := make(map[string]float64)
	for _, indVar := range predictionData.IndependentVars {
		varVals[indVar.Name] = indVar.Value
	}

	for _, varName := range varNames {
		coeff, ok := coeffs[varName]
		if !ok {
			return fmt.Errorf("Could not find model coefficient %s", varName)
		}

		val, ok := varVals[varName]
		if !ok {
			return fmt.Errorf("Expected value for variable %s", varName)
		}

		prediction = prediction + coeff * val
	}

	predictionData.Prediction = prediction

	return nil
}
