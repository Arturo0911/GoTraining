package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const pathFile = "log_diff_series.csv"
const passFile = "../AirPassengers.csv"

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
	passengerValues := passDF.Col("log_differenced_passengers").Float()

	// Calculate the coefficients.
	coeffs, intercept := autoregressive(passengerValues, 2)

	fmt.Printf("\nlog(x(t)) - log(x(t-1)) = %0.6f+lag1*%0.6f + lag2*%0.6f\n\n", intercept, coeffs[0], coeffs[1])

}

func maeError() {
	transFile, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
	}
	defer transFile.Close()

	transReader := csv.NewReader(transFile)
	transReader.FieldsPerRecord = 2
	transdata, err := transReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Looping over the data predicting the transformed
	var transPredictions []float64
	for i := range transdata {

		// skip the header and the first two observations
		// because we need two lags to make predictions
		if i == 0 || i == 1 || i == 2 {
			continue
		}

		// Parse the firs lag.
		lagOne, err := strconv.ParseFloat(transdata[i-1][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the second lag.
		lagTwo, err := strconv.ParseFloat(transdata[i-2][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict the transformed variable with our trained AR model

		transPredictions = append(transPredictions,
			0.008159+0.234953*lagOne-0.173682*lagTwo)
	}

	// Open the original dataset file.
	origFile, err := os.Open(passFile)
	if err != nil {
		log.Fatal(err)
	}
	defer origFile.Close()

	// Create a csv reader reading from the opened file
	origReader := csv.NewReader(origFile)
	origReader.FieldsPerRecord = 2

	origData, err := origReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// pts* will hold the values for plotting
	ptsObs := make(plotter.XYs, len(transPredictions))
	ptsPred := make(plotter.XYs, len(transPredictions))

	// Reverse the transformation and calculate the MAE
	var mAE float64
	var cumSum float64
	fmt.Println(len(origData) - 1)
	for i := 4; i <= len(origData)-1; i++ {

		// Parse the original observation.
		observed, err := strconv.ParseFloat(origData[i][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Parse the original date.
		date, err := strconv.ParseFloat(origData[i][0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Get the cumulative sum up to the index in
		// the transformed predictions.
		cumSum += transPredictions[i-4]

		// Calculate the reverse transformed prediction.
		predicted := math.Exp(math.Log(observed) + cumSum)

		// Accumulate the MAE.
		mAE += math.Abs(observed-predicted) / float64(len(transPredictions))

		// Fill in the points for plotting.
		ptsObs[i-4].X = date
		ptsPred[i-4].X = date
		ptsObs[i-4].Y = observed
		ptsPred[i-4].Y = predicted
	}
	// printing the MAE
	fmt.Printf("\nMAE = %0.2f\n\n", mAE)

	p := plot.New()
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

	// Save the plot to a png file
	p.Add(lObs, lPred)
	p.Legend.Add("Observed", lObs)
	p.Legend.Add("Predicted", lPred)
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "passengers_ts.png"); err != nil {
		log.Fatal(err)
	}

}

func main() {

	readerFile()
	maeError()
}
