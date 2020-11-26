package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

type neuralNetConfig struct {
	inputNeurons  int
	outputNeurons int
	hiddenNeurons int
	numEpochs     int
	learningRate  float64
}

type neuralNet struct {
	config  neuralNetConfig
	wHidden *mat.Dense
	bHidden *mat.Dense
	wOut    *mat.Dense
	bOut    *mat.Dense
}

var (
	trainingName = "iris_training.csv"
	testName = "iris_training.csv"
	trainingPath = filepath.Join(os.Getenv("MLGO"), "data", trainingName)
	testPath = filepath.Join(os.Getenv("MLGO"), "data", testName)

	config = neuralNetConfig{
		inputNeurons:  4,
		outputNeurons: 3,
		hiddenNeurons: 3,
		numEpochs:     5000,
		learningRate:  0.3,
	}
)

func main() {
	// training data
	f, err := os.Open(trainingPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 7

	trainData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	inputsData := make([]float64, 4 * len(trainData))
	labelsData := make([]float64, 3 * len(trainData))

	var inputsIndex int
	var labelsIndex int

	for idx, record := range trainData {
		if idx == 0 {
			continue
		}

		for i, val := range record {
			parsedVal, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}

			if i == 4 || i == 5 || i == 6 {
				labelsData[labelsIndex] = parsedVal
				labelsIndex++
				continue
			}

			inputsData[inputsIndex] = parsedVal
			inputsIndex++
		}
	}

	inputs := mat.NewDense(len(trainData), 4, inputsData)
	labels := mat.NewDense(len(trainData), 3, labelsData)

	// train
	network := newNetwork(config)
	if err := network.train(inputs, labels); err != nil {
		log.Fatal(err)
	}

	// test data
	f, err = os.Open(testPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader = csv.NewReader(f)
	reader.FieldsPerRecord = 7

	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	inputsData = make([]float64, 4 * len(testData))
	labelsData = make([]float64, 3 * len(testData))

	inputsIndex = 0
	labelsIndex = 0

	for idx, record := range testData {
		if idx == 0 {
			continue
		}

		for i, val := range record {
			parsedVal, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}

			if i == 4 || i == 5 || i == 6 {
				labelsData[labelsIndex] = parsedVal
				labelsIndex++
				continue
			}

			inputsData[inputsIndex] = parsedVal
			inputsIndex++
		}
	}

	testInputs := mat.NewDense(len(testData), 4, inputsData)
	testLabels := mat.NewDense(len(testData), 3, labelsData)

	// prediction
	predictions, err := network.predict(testInputs)
	if err != nil {
		log.Fatal(err)
	}

	var truePosNeg int

	numPreds, _ := predictions.Dims()

	for i := 0; i < numPreds; i++ {
		labelRow := mat.Row(nil, i, testLabels)

		var species int

		for idx, label := range labelRow {
			if label == 1.0 {
				species = idx
				break
			}
		}

		if predictions.At(i, species) == floats.Max(mat.Row(nil, i, predictions)) {
			truePosNeg++
		}
	}

	accuracy := float64(truePosNeg) / float64(numPreds)
	fmt.Printf("\nAccuracy = %.2f\n\n", accuracy)
}

func newNetwork(config neuralNetConfig) *neuralNet {
	return &neuralNet{config: config}
}

func (nn *neuralNet) train(x, y *mat.Dense) error {
	// reproducibility
	s := rand.NewSource(42)
	r := rand.New(s)

	wHiddenRaw := make([]float64, nn.config.hiddenNeurons * nn.config.inputNeurons)
	bHiddenRaw := make([]float64, nn.config.hiddenNeurons)
	wOutRaw := make([]float64, nn.config.outputNeurons * nn.config.hiddenNeurons)
	bOutRaw := make([]float64, nn.config.outputNeurons)

	for _, param := range [][]float64{wHiddenRaw, bHiddenRaw, wOutRaw, bOutRaw} {
		for i:= range param{
			param[i] = r.Float64()
		}
	}

	numData, _ := x.Dims()

	wHidden := mat.NewDense(nn.config.inputNeurons, nn.config.hiddenNeurons, wHiddenRaw) // (4 x 3)
	bHidden := mat.NewDense(1, nn.config.hiddenNeurons, bHiddenRaw) // (1 x 3)
	wOut := mat.NewDense(nn.config.hiddenNeurons, nn.config.outputNeurons, wOutRaw) // (3 x 3)
	bOut := mat.NewDense(1, nn.config.outputNeurons, bOutRaw) // (1 x 3)

	output := mat.NewDense(numData, 3, nil)

	for i := 0; i < nn.config.numEpochs; i++ {
		hiddenLayerInput := mat.NewDense(numData, 3, nil)
		hiddenLayerInput.Mul(x, wHidden) // (numData x 4) mul (4 x 3) = (numData x 3)
		addBHidden := func(_, col int, v float64) float64 { return v + bHidden.At(0, col) }
		hiddenLayerInput.Apply(addBHidden, hiddenLayerInput)

		hiddenLayerActivations := mat.NewDense(numData, 3, nil)
		applySigmoid := func(_, _ int, v float64) float64 { return sigmoid(v) }
		hiddenLayerActivations.Apply(applySigmoid, hiddenLayerInput)

		outputLayerInput := mat.NewDense(numData, 3, nil)
		outputLayerInput.Mul(hiddenLayerActivations, wOut) // (numData x 3) mul (3 x 3) = (numData x 3)
		addBOut := func(_, col int, v float64) float64 { return v + bOut.At(0, col) }
		outputLayerInput.Apply(addBOut, outputLayerInput)
		output.Apply(applySigmoid, outputLayerInput)

		networkError := mat.NewDense(numData, 3, nil)
		networkError.Sub(y, output)

		slopeOutputLayer := mat.NewDense(numData, 3, nil)
		applySigmoidPrime := func(_, _ int, v float64) float64 { return sigmoidPrime(v) }
		slopeOutputLayer.Apply(applySigmoidPrime, output)
		slopeHiddenLayer := mat.NewDense(numData, 3, nil)
		slopeHiddenLayer.Apply(applySigmoidPrime, hiddenLayerActivations)

		dOutput := mat.NewDense(numData, 3, nil)
		dOutput.MulElem(networkError, slopeOutputLayer)
		errorAtHiddenLayer := mat.NewDense(numData, 3, nil)
		errorAtHiddenLayer.Mul(dOutput, wOut.T()) // (numData x 3) mul (3 x 3) = (numData x 3)

		dHiddenLayer := mat.NewDense(numData, 3, nil)
		dHiddenLayer.MulElem(errorAtHiddenLayer, slopeHiddenLayer)

		wOutAdj := mat.NewDense(3, 3, nil)
		wOutAdj.Mul(hiddenLayerActivations.T(), dOutput) // (3 x numData) mul (numData x 3) = (3 x 3)
		wOutAdj.Scale(nn.config.learningRate, wOutAdj)
		wOut.Add(wOut, wOutAdj)

		bOutAdj, err := sumAlongAxis(0, dOutput)
		if err != nil {
			return err
		}
		bOutAdj.Scale(nn.config.learningRate, bOutAdj)
		bOut.Add(bOut, bOutAdj)

		wHiddenAdj := mat.NewDense(4, 3, nil)
		wHiddenAdj.Mul(x.T(), dHiddenLayer) // (4 x numData) mul (numData x 3) = (4 x 3)
		wHiddenAdj.Scale(nn.config.learningRate, wHiddenAdj)
		wHidden.Add(wHidden, wHiddenAdj)

		bHiddenAdj, err := sumAlongAxis(0, dHiddenLayer)
		if err != nil {
			return err
		}
		bHiddenAdj.Scale(nn.config.learningRate, bHiddenAdj)
		bHidden.Add(bHidden, bHiddenAdj)
	}

	nn.wHidden = wHidden
	nn.bHidden = bHidden
	nn.wOut = wOut
	nn.bOut = bOut

	return nil
}

func (nn *neuralNet) predict(x *mat.Dense) (*mat.Dense, error) {
	if nn.wHidden == nil || nn.wOut == nil || nn.bHidden == nil || nn.bOut == nil {
		return nil, errors.New("the supplied neural net weights and biases are empty")
	}

	numData, _ := x.Dims()

	output := mat.NewDense(numData, 3, nil)

	hiddenLayerInput := mat.NewDense(numData, 3, nil)
	hiddenLayerInput.Mul(x, nn.wHidden) // (numData x 4) mul (4 x 3) = (numData x 3)
	addBHidden := func(_, col int, v float64) float64 { return v + nn.bHidden.At(0, col) }
	hiddenLayerInput.Apply(addBHidden, hiddenLayerInput)

	hiddenLayerActivations := mat.NewDense(numData, 3, nil)
	applySigmoid := func(_, _ int, v float64) float64 { return sigmoid(v) }
	hiddenLayerActivations.Apply(applySigmoid, hiddenLayerInput)

	outputLayerInput := mat.NewDense(numData, 3, nil)
	outputLayerInput.Mul(hiddenLayerActivations, nn.wOut) // (numData x 3) mul (3 x 3) = (numData x 3)
	addBOut := func(_, col int, v float64) float64 { return v + nn.bOut.At(0, col) }
	outputLayerInput.Apply(addBOut, outputLayerInput)
	output.Apply(applySigmoid, outputLayerInput)

	return output, nil
}

func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func sigmoidPrime(x float64) float64 {
	return x * (1.0 - x)
}

func sumAlongAxis(axis int, m *mat.Dense) (*mat.Dense, error) {
	numRows, numCols := m.Dims()

	var output *mat.Dense

	switch axis {
	case 0:
		data := make([]float64, numCols)
		for i := 0; i < numCols; i++ {
			col := mat.Col(nil, i, m)
			data[i] = floats.Sum(col)
		}
		output = mat.NewDense(1, numCols, data)
	case 1:
		data := make([]float64, numRows)
		for i := 0; i < numRows; i++ {
			row := mat.Row(nil, i, m)
			data[i] = floats.Sum(row)
		}
		output = mat.NewDense(numRows, 1, data)
	default:
		return nil, errors.New("invalid axis, must be 0 ro 1")
	}

	return output, nil
}
