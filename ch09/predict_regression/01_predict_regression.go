package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

	attrFiles := []string{"1.json", "2.json", "3.json"}
	for _, fileName := range attrFiles {
		inVarPath := filepath.Join(*inVarDirPtr, fileName)
		f, err := ioutil.ReadFile(inVarPath)
		if err != nil {
			log.Fatal(err)
		}

		var predictionData PredictionData
		if err := json.Unmarshal(f, &predictionData); err != nil {
			log.Fatal(err)
		}

		if err := Predict(&modelInfo, &predictionData); err != nil {
			log.Fatal(err)
		}

		outputData, err := json.MarshalIndent(predictionData, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		resFileName := "predicted_" + fileName
		if err := ioutil.WriteFile(filepath.Join(*outDirPtr, resFileName), outputData, 0644); err != nil {
			log.Fatal(err)
		}
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
