package main

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot/plotter"
)

// Data visualitations to quantify how distributions work!

func GraphHistogram(pathFile string) {

	irisCsv, err := os.Open(pathFile)

	if err != nil {
		log.Fatal(err)
	}

	defer irisCsv.Close()

	IrisDF := dataframe.ReadCSV(irisCsv)

	// Now create a histogrma for each of the feature columns in the
	// dataset.

	for _, colName := range IrisDF.Names() {

		// If the column is one of the feature columns, let's
		// create  a histogram of the values

		if conName != "species" {
			// Create a plotter.Values value and fill it with the
			// values from the respective column of the dataframe

			v := make(plotter.Values, IrisDF)
		}
	}
}

func main() {

	pathFile := "../iris_labeled.csv"

	GraphHistogram(pathFile)
}
