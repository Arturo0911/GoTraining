package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/stat"
)

const pathFile = "../AirPassengers.csv"

// Understanding time series jargon

func ACF(x []float64, lag int) float64 {

	// Shift the series
	xAdj := x[lag:len(x)]
	xLag := x[0 : len(x)*-lag]

	// Numerator will hold our accumulated numerator, and
	// denominator will hold our accumulated denominator.
	var numerator float64
	var denominator float64

	// Calculate the mean of our x values, which will be used
	// in each term of the autocorrelation
	xBar := stat.Mean(x, nil)

	// Calculate the numerator.
	for idx, xVal := range xAdj {
		numerator += ((xVal - xBar) * (xLag[idx] - xBar))
	}

	// Calculate denominator
	for _, xVal := range x {
		denominator += math.Pow(xVal-xBar, 2)
	}

	return numerator / denominator
}

func TimeSeries() {

	// Open the csv file.
	passengerFile, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
	}
	defer passengerFile.Close()

	passDF := dataframe.ReadCSV(passengerFile)

	passengers := passDF.Col("AirPassengers").Float()
	//fmt.Println(passengers)
	// Loop over varios values of lag in the series.
	fmt.Println("Autocorrelation: ")
	for i := 1; i < 11; i++ {

		// Shift the series

		adjusted := passengers[i:len(passengers)]
		lag := passengers[0 : len(passengers)-i]

		ac := stat.Correlation(adjusted, lag, nil)
		fmt.Printf("\nLag %d period: %0.2f\n", i, ac)
	}
}

func main() {
	TimeSeries()
}
