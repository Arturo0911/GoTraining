package main

import (
	"strconv"

	"github.com/sajari/regression"
)

func autoregressive(x []float64, lag int) ([]float64, float64) {

	// Crete a regreession.Regression value needed to train
	// a model using sajari repository

	var r regression.Regression
	r.SetObserved("x")

	// Define the current lag and all fo the intermediate lags.
	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}

	// Shift the series
	xAdj := x[lag:len(x)]
}

func main() {

}
