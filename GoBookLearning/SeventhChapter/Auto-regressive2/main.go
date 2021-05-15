package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
)

const pathFile = "log_diff_series.csv"

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
	xAdj := x[lag:]

	// loop over the series creating the dataset
	// for the regression.
	for i, xVal := range xAdj {

		// Loop over the intermediate lags to build up
		// our independend variables.
		laggedVariables := make([]float64, lag)
		for idx := 1; idx <= lag; idx++ {

			// Get the lagged series variables.
			laggedVariables[idx-1] = x[lag+i-idx]
		}

		// Add these points to the regression value.
		r.Train(regression.DataPoint(xVal, laggedVariables))
	}

	// Fit the regression
	r.Run()

	// Coeff hold the coefficients for our lags.
	var coeff []float64
	for i := 1; i <= lag; i++ {
		coeff = append(coeff, r.Coeff(i))
	}
	return coeff, r.Coeff(0)
}

func readerFile() {
	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passDF := dataframe.ReadCSV(file)
	passengerValues := passDF.Col("differenced_passengers").Float()

	// Calculate the coefficients.
	coeffs, intercept := autoregressive(passengerValues, 2)

	fmt.Printf("\nlog(x(t)) - log(x(t-1)) = %0.6f+lag1*%0.6f + lag2*%0.6f\n\n", intercept, coeffs[0], coeffs[1])

}

func main() {

	readerFile()
}
