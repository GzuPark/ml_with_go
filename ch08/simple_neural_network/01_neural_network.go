package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"

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
	input_data = []float64{
		1.0, 0.0, 1.0, 0.0,
		1.0, 0.0, 1.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
	}
	labels_data = []float64{1.0, 1.0, 0.0}
	config = neuralNetConfig{
		inputNeurons:  4,
		outputNeurons: 1,
		hiddenNeurons: 3,
		numEpochs:     5000,
		learningRate:  0.3,
	}
)

func main() {
	input := mat.NewDense(3, 4, input_data)
	labels := mat.NewDense(3, 1, labels_data)

	network := newNetwork(config)
	if err := network.train(input, labels); err != nil {
		log.Fatal(err)
	}

	f := mat.Formatted(network.wHidden, mat.Prefix("          "))
	fmt.Printf("\nwHidden = %v\n\n", f)

	f = mat.Formatted(network.bHidden, mat.Prefix("    "))
	fmt.Printf("\nbHidden = %v\n\n", f)

	f = mat.Formatted(network.wOut, mat.Prefix("       "))
	fmt.Printf("\nwOut = %v\n\n", f)

	f = mat.Formatted(network.bOut, mat.Prefix("    "))
	fmt.Printf("\nbOut = %v\n\n", f)
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

	wHidden := mat.NewDense(nn.config.inputNeurons, nn.config.hiddenNeurons, wHiddenRaw) // (4 x 3)
	bHidden := mat.NewDense(1, nn.config.hiddenNeurons, bHiddenRaw) // (1 x 3)
	wOut := mat.NewDense(nn.config.hiddenNeurons, nn.config.outputNeurons, wOutRaw) // (3 x 1)
	bOut := mat.NewDense(1, nn.config.outputNeurons, bOutRaw) // (1 x 1)

	output := mat.NewDense(3, 1, nil)

	for i := 0; i < nn.config.numEpochs; i++ {
		hiddenLayerInput := mat.NewDense(3, 3, nil)
		hiddenLayerInput.Mul(x, wHidden) // (3 x 4) mul (4 x 3) = (3 x 3)
		addBHidden := func(_, col int, v float64) float64 { return v + bHidden.At(0, col) }
		hiddenLayerInput.Apply(addBHidden, hiddenLayerInput)

		hiddenLayerActivations := mat.NewDense(3, 3, nil)
		applySigmoid := func(_, _ int, v float64) float64 { return sigmoid(v) }
		hiddenLayerActivations.Apply(applySigmoid, hiddenLayerInput)

		outputLayerInput := mat.NewDense(3, 1, nil)
		outputLayerInput.Mul(hiddenLayerActivations, wOut) // (3 x 3) mul (3 x 1) = (3 x 1)
		addBOut := func(_, col int, v float64) float64 { return v + bOut.At(0, col) }
		outputLayerInput.Apply(addBOut, outputLayerInput)
		output.Apply(applySigmoid, outputLayerInput)

		networkError := mat.NewDense(3, 1, nil)
		networkError.Sub(y, output)

		slopeOutputLayer := mat.NewDense(3, 1, nil)
		applySigmoidPrime := func(_, _ int, v float64) float64 { return sigmoidPrime(v) }
		slopeOutputLayer.Apply(applySigmoidPrime, output)
		slopeHiddenLayer := mat.NewDense(3, 3, nil)
		slopeHiddenLayer.Apply(applySigmoidPrime, hiddenLayerActivations)

		dOutput := mat.NewDense(3, 1, nil)
		dOutput.MulElem(networkError, slopeOutputLayer)
		errorAtHiddenLayer := mat.NewDense(3, 3, nil)
		errorAtHiddenLayer.Mul(dOutput, wOut.T()) // (3 x 1) mul (1 x 3) = (3 x 3)

		dHiddenLayer := mat.NewDense(3, 3, nil)
		dHiddenLayer.MulElem(errorAtHiddenLayer, slopeHiddenLayer)

		wOutAdj := mat.NewDense(3, 1, nil)
		wOutAdj.Mul(hiddenLayerActivations.T(), dOutput) // (3 x 3) mul (3 x 1) = (3 x 1)
		wOutAdj.Scale(nn.config.learningRate, wOutAdj)
		wOut.Add(wOut, wOutAdj)

		bOutAdj, err := sumAlongAxis(0, dOutput)
		if err != nil {
			return err
		}
		bOutAdj.Scale(nn.config.learningRate, bOutAdj)
		bOut.Add(bOut, bOutAdj)

		wHiddenAdj := mat.NewDense(4, 3, nil)
		wHiddenAdj.Mul(x.T(), dHiddenLayer) // (4 x 3) mul (3 x 3) = (4 x 3)
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
