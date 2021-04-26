package main

import (
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
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
	}
}

func main() {

	pathFile := "../iris_labeled.csv"

	GraphHistogram(pathFile)
}
