package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// NeuralNet contains all of the information
// that defines a trained neural network
type neuralNet struct {
	config  neuralNetConfig
	wHidden *mat.Dense
	bHidden *mat.Dense
	wOut    *mat.Dense
	bOut    *mat.Dense
}

// neuralNetConfig defines our neural network
// architecture and learning parameters.
type neuralNetConfig struct {
	inputNeurons  int
	outputNeurons int
	hiddenNeurons int
	numEpochs     int
	learningRate  float64
}

// Sigmoid implements the sigmoid function
// for use in activation functions.
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// SigmoidPrime implements the derivative
// of the sigmoid function for backPropagation
func sigmoidPrime(x float64) float64 {
	return x * (1.0 - x)
}

// NewNetwork initilizes a new neural network.
func newNetwork(config neuralNetConfig) *neuralNet {
	return &neuralNet{config: config}
}

// Train traings a neural network using backPropagation
func (nn *neuralNet) train(x, y *mat.Dense) error {

	// Initialize biases/weights.
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	wHiddenRaw := make([]float64,
		nn.config.hiddenNeurons*nn.config.inputNeurons)
	bHiddenRaw := make([]float64, nn.config.hiddenNeurons)
	wOutRaw := make([]float64,
		nn.config.outputNeurons*nn.config.hiddenNeurons)
	bOutRaw := make([]float64, nn.config.outputNeurons)

	for _, param := range [][]float64{wHiddenRaw, bHiddenRaw, wOutRaw, bOutRaw} {

		for i := range param {
			param[i] = randGen.Float64()
		}
	}

	wHidden := mat.NewDense(nn.config.inputNeurons, nn.config.hiddenNeurons, wHiddenRaw)
	bHidden := mat.NewDense(1, nn.config.hiddenNeurons, bHiddenRaw)
	wOut := mat.NewDense(nn.config.hiddenNeurons, nn.config.outputNeurons, wOutRaw)
	bOut := mat.NewDense(1, nn.config.outputNeurons, bOutRaw)

	// Define the output of the neural network.
	var output mat.Dense
	//fmt.Println(output)

	// Loop over the number of epochs utilizing
	// backpropagration to train our model.
	for i := 0; i < nn.config.numEpochs; i++ {

		var hiddenLayerInput mat.Dense
		//fmt.Println(hiddenLayerInput)
		hiddenLayerInput.Mul(x, wHidden)
		addBHidden := func(_, col int,
			v float64) float64 {
			return v + bHidden.At(0, col)
		}
		hiddenLayerInput.Apply(addBHidden, &hiddenLayerInput)
		var hiddenLayerActivations mat.Dense
		applySigmoid := func(_, _ int, v float64) float64 {
			return sigmoid(v)
		}
		hiddenLayerActivations.Apply(applySigmoid,
			&hiddenLayerInput)

		var outputLayerInput mat.Dense
		outputLayerInput.Mul(&hiddenLayerActivations, wOut)
		addBOut := func(_, col int, v float64) float64 {
			return v + bOut.At(0, col)
		}
		outputLayerInput.Apply(addBOut, &outputLayerInput)
		output.Apply(applySigmoid, &outputLayerInput)

		//	 _                _
		//	| |              | |
		//	| |__   __ _  ___| | __
		//	| '_ \ / _` |/ __| |/ /
		//	| |_) | (_| | (__|   <
		//	|_.__/ \__,_|\___|_|\_\
		//										     _    _
		//										    | |  (_)
		//	 _ __  _ __ ___  _ __   __ _  __ _  __ _| |_ _  ___  _ __
		//	| '_ \| '__/ _ \| '_ \ / _` |/ _` |/ _` | __| |/ _ \| '_ \
		//	| |_) | | | (_) | |_) | (_| | (_| | (_| | |_| | (_) | | | |
		//	| .__/|_|  \___/| .__/ \__,_|\__, |\__,_|\__|_|\___/|_| |_|
		//	| |             | |           __/ |
		//	|_|             |_|          |___/

		var networkError mat.Dense
		networkError.Sub(y, &output)

		var slopeOutputLayer mat.Dense
		applySigmoidPrime := func(_, _ int, v float64) float64 {
			return sigmoidPrime(v)
		}
		slopeOutputLayer.Apply(applySigmoidPrime, &output)
		var slopeHiddenLayer mat.Dense
		slopeHiddenLayer.Apply(applySigmoidPrime, &hiddenLayerActivations)

		var dOutput mat.Dense
		dOutput.MulElem(&networkError, &slopeOutputLayer)
		var errorAtHiddenLayer mat.Dense
		errorAtHiddenLayer.Mul(&dOutput, wOut.T())

		var dHiddenLayer mat.Dense
		dHiddenLayer.MulElem(&errorAtHiddenLayer, &slopeHiddenLayer)

		//				 _ _           _
		//		/\      | (_)         | |
		//	   /  \   __| |_ _   _ ___| |_
		//	  / /\ \ / _` | | | | / __| __|
		//	 / ____ \ (_| | | |_| \__ \ |_
		//	/_/    \_\__,_| |\__,_|___/\__|
		//				 _/ |
		//				|__/
		//										 _
		//										| |
		//	_ __   __ _ _ __ __ _ _ __ ___   ___| |_ ___ _ __ ___
		//	| '_ \ / _` | '__/ _` | '_ ` _ \ / _ \ __/ _ \ '__/ __|
		//	| |_) | (_| | | | (_| | | | | | |  __/ ||  __/ |  \__ \
		//	| .__/ \__,_|_|  \__,_|_| |_| |_|\___|\__\___|_|  |___/
		//	| |
		//	|_|

		var wOudAdj mat.Dense
		wOudAdj.Mul(hiddenLayerActivations.T(), &dOutput)
		wOudAdj.Scale(nn.config.learningRate, &wOudAdj)
		wOut.Add(wOut, &wOudAdj)

		bOutAdj, err := sumAlongAxis(0, &dOutput)
		if err != nil {
			log.Fatal(err)
		}
		bOutAdj.Scale(nn.config.learningRate, bOutAdj)
		bOutAdj.Add(bOut, bOutAdj)

		var wHiddenAdj mat.Dense
		wHiddenAdj.Mul(x.T(), &dHiddenLayer)
		wHiddenAdj.Scale(nn.config.learningRate, &wHiddenAdj)
		wHidden.Add(wHidden, &wHiddenAdj)

		bHiddenAdj, err := sumAlongAxis(0, &dHiddenLayer)
		if err != nil {
			return err
		}
		//fmt.Println(bHiddenAdj)
		bHiddenAdj.Scale(nn.config.learningRate, bHiddenAdj)
		bHidden.Add(bHidden, bHiddenAdj)

	}

	nn.wHidden = wHidden
	nn.bHidden = bHidden
	nn.wOut = wOut
	nn.bOut = bOut

	return nil
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
		return nil, errors.New("invalid axis, must be 0 or 1")
	}

	return output, nil
}

func main() {

	// An empty matrix is one that has zero size.
	// Empty matrices are used to allow the destination
	// of a matrix operation to assume the correct size
	// automatically. This operation will re-use the
	// backing data, if available, or will allocate new data
	// if necessary. The IsEmpty method returns whether the
	// given matrix is empty. The zero-value of a matrix is
	// empty, and is useful for easily getting the result of
	// matrix operations.

	/*
		var c mat.Dense // construct a new zero-value matrix

		c.Mul(a, a) // c is automatically adjusted to be the right size

	*/

	// Define our input attributes.
	input := mat.NewDense(3, 4, []float64{
		1.0, 0.0, 1.0, 0.0,
		1.0, 0.0, 1.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
	})
	fmt.Println(input.Dims())
	fmt.Println(input)
	fmt.Println(input.Caps())

	// Define our labels
	labels := mat.NewDense(3, 1, []float64{1.0, 1.0, 0.0})

	// Define our network architecture and
	// learning parameters.
	config := neuralNetConfig{
		inputNeurons:  4,
		outputNeurons: 1,
		hiddenNeurons: 3,
		numEpochs:     5000,
		learningRate:  0.3,
	}

	// Train the nwueal network.
	network := newNetwork(config)
	if err := network.train(input, labels); err != nil {
		log.Fatal(err)
	}

	// output the weights that define our network!
	f := mat.Formatted(network.wHidden, mat.Prefix(" "))
	fmt.Printf("\nwHidden = %v \n\n", f)

	f = mat.Formatted(network.bHidden, mat.Prefix(" "))
	fmt.Printf("\nbHidden = %v \n\n", f)

	f = mat.Formatted(network.wOut, mat.Prefix(" "))
	fmt.Printf("\nwOut = %v \n\n", f)

	f = mat.Formatted(network.bOut, mat.Prefix(" "))
	fmt.Printf("\nbOut = %v \n\n", f)

	//fmt.Println(input)

}
