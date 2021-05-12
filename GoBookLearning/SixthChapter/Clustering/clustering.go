package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/gonum/floats"
)

type centroid []float64

func main() {

	// pull in the csv file
	irisFile, err := os.Open("../datasets/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// create a dataframe
	irisDF := dataframe.ReadCSV(irisFile)
	fmt.Println(irisDF)
	// Define the names of the hree separate species
	// contained in the csv file
	speciesNames := []string{
		"Iris-setosa",
		"Iris-versicolor",
		"Iris-virginica",
	}

	// Crate a map to hold our centroid information
	centroids := make(map[string]centroid)

	// Filter the dataset into three separate dataframes,
	// each corresponding to ne of the Iris species.

	clusters := make(map[string]dataframe.DataFrame)
	for _, species := range speciesNames {

		// Filter the original dataset.
		filter := dataframe.F{
			Colname:    "species",
			Comparator: "==",
			Comparando: species,
		}
		filtered := irisDF.Filter(filter)
		//fmt.Println(filtered)
		// Calculate the mean of features
		summaryDF := filtered.Describe()
		//fmt.Println(summaryDF.Names())
		// put each dimension's mean into the corresponding centroid.

		var c centroid

		for _, feature := range summaryDF.Names() {

			// skip the irrelevant columns
			if feature == "column" || feature == "species" {
				continue
			}
			c = append(c, summaryDF.Col(feature).Float()[0])
		}

		// Add this centroid to our map
		centroids[species] = c
		clusters[species] = filtered
	}

	// As a sanity check, output oour centroids
	for _, species := range speciesNames {
		fmt.Printf("%s centroid: %v\n", species, centroids[species])
		fmt.Printf("\n")
	}

	// Covert our labels into a slice of strings and create a slice
	// of float column names for convenience.
	lables := irisDF.Col("species").Records()
	floatColumns := []string{
		"sepal_length",
		"sepal_width",
		"petal_length",
		"petal_width",
	}

	// Loop over the records accumukating the average silhoutte
	// ccoefficient.
	var silhoutte float64

	for idx, label := range lables {
		// a will store our accumulated value for a.
		var a float64

		// loop over the data points in the same cluser.

		for i := 0; i < clusters[label].Nrow(); i++ {
			// Get the data point for comparation
			current := dfFloatRow(irisDF, floatColumns, idx)
			other := dfFloatRow(clusters[label], floatColumns, i)

			// Add to a
			a += floats.Distance(current, other, 2) / float64(clusters[label].Nrow())

		}
		// Determine the neares other clluster
	}

}

func dfFloatRow(df dataframe.DataFrame, names []string, idx int) []float64 {

	var row []float64
	for _, name := range names {
		row = append(row, df.Col(name).Float()[idx])
	}
	return row

}
