package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

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
	boutRaw := make([]float64, nn.config.outputNeurons)

	for _, param := range [][]float64{wHiddenRaw,
		bHiddenRaw, wOutRaw, boutRaw} {

		for i := range param {
			param[i] = randGen.Float64()
			fmt.Println(param[i])
		}
	}

	wHidden := mat.NewDense(nn.config.inputNeurons,
		nn.config.hiddenNeurons, wHiddenRaw)
	bHidden := mat.NewDense(1, nn.config.hiddenNeurons,
		bHiddenRaw)
	wOut := mat.NewDense(nn.config.hiddenNeurons,
		nn.config.outputNeurons, wOutRaw)
	bOut := mat.NewDense(1, nn.config.outputNeurons, boutRaw)

	// Define the output of the neural network.
	output := mat.NewDense(0, 0, nil)

	// Loop over the number of epochs utilizing
	// backpropagration to train our model.
	for i := 0; i < nn.config.numEpochs; i++ {

		hiddenLayerInput := mat.NewDense(0, 0, nil)
		hiddenLayerInput.Mul(x, wHidden)
		addBHidden := func(_, col int,
			v float64) float64 {
			return v + bHidden.At(0, col)
		}
		hiddenLayerInput.Apply(addBHidden, hiddenLayerInput)
		hiddenLayerActivations := mat.NewDense(0, 0, nil)
		applySigmoid := func(_, _ int, v float64) float64 {
			return sigmoid(v)
		}
		hiddenLayerActivations.Apply(applySigmoid,
			hiddenLayerInput)

		outputLayerInput := mat.NewDense(0, 0, nil)
		outputLayerInput.Mul(hiddenLayerActivations, wOut)
		addBOut := func(_, col int, v float64) float64 {
			return v + bOut.At(0, col)
		}
		outputLayerInput.Apply(addBOut, outputLayerInput)
		output.Apply(applySigmoid, outputLayerInput)

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

		networkError := mat.NewDense(0, 0, nil)
		networkError.Sub(y, output)

		slopeOutputLayer := mat.NewDense(0, 0, nil)
		applySigmoidPrime := func(_, _ int, v float64) float64 {
			return sigmoidPrime(v)
		}
		slopeOutputLayer.Apply(applySigmoidPrime, output)
		slopeHiddenLayer := mat.NewDense(0, 0, nil)
		slopeHiddenLayer.Apply(applySigmoidPrime, hiddenLayerActivations)

		dOutput := mat.NewDense(0, 0, nil)
		dOutput.MulElem(networkError, slopeOutputLayer)
		errorAtHiddenLayer := mat.NewDense(0, 0, nil)
		errorAtHiddenLayer.Mul(dOutput, wOut.T())

		dHiddenLayer := mat.NewDense(0, 0, nil)
		dHiddenLayer.MulElem(errorAtHiddenLayer, slopeHiddenLayer)

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

		wOudAdj := mat.NewDense(0, 0, nil)
		wOudAdj.Mul(hiddenLayerActivations.T(), dOutput)
		wOudAdj.Scale(nn.config.learningRate, wOudAdj)
		wOut.Add(wOut, wOudAdj)

		bOutAdj, err := sumAlongAxis(0, dOutput)
		if err != nil {
			log.Fatal(err)
		}
		bOutAdj.Scale(nn.config.learningRate, bOutAdj)
		bOutAdj.Add(bOut, bOutAdj)

	}

	return nil
}

func sumAlongAxis(axis int, m *mat.Dense) (*mat.Dense, error) {}

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

	var value *mat.Dense
	fmt.Print(value) //

}
