package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Seting the Y variable as dependent variable
func CheckingCorrelation(pathFile string) {

	file, err := os.Open(pathFile)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	dataCSV := dataframe.ReadCSV(file)

	// Extract the target of the column
	// chosse the column, and the name
	yVals := dataCSV.Col("Sales").Float()
	fmt.Println(yVals)
	// create a scatter a polot for each of the features in the
	// dataset

	for _, colName := range dataCSV.Names() {
		// pts will hold the value for plotting

		pts := make(plotter.XYs, dataCSV.Nrow())

		// fill pts with data.

		for i, floatVal := range dataCSV.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}

		// create the plot

		p := plot.New()

		p.X.Label.Text = colName
		p.Y.Label.Text = "y"

		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)

		if err != nil {
			log.Fatal(err)
		}

		s.GlyphStyle.Radius = vg.Points(3)
		p.Add(s)
		// save the plot to a PNG file

		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}

}

// Making plot charts
func ReadingAdvertising(pathFile string) {

	file, err := os.Open(pathFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	advertDF := dataframe.ReadCSV(file)

	/*// Use the describe metjod to calculate
	// summary statistics for all the columns in one shot

	adverSummary := advertDF.Describe()

	fmt.Println(adverSummary)*/

	// Extract the target column.
	//yVals := advertDF.Col("Sales").Float()

	// create a histogram for each of the columns in the dataset.
	for _, colName := range advertDF.Names() { // Printing the headers
		// create a plotter. values value and fill it with the
		// values from the respective column of the dataframe.

		plotVals := make(plotter.Values, advertDF.Nrow())

		for i, floatVal := range advertDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Making a plot and set its title

		p := plot.New()
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		// Create a histogram of our values drawn
		// from the standard normal.

		h, err := plotter.NewHist(plotVals, 16)

		if err != nil {
			log.Fatal(err)
		}

		// Normalizing the data.
		h.Normalize(1)

		// Add the histogram to the plot.

		p.Add(h)

		// saving the plot to a PNG file.

		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}

}

// Setting the data to be training and testing
func CreatingDataTraining(pathFile string) {

	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	advertDF := dataframe.ReadCSV(file)

	// Calculating, number of elements in each set

	traningNum := (4 * advertDF.Nrow()) / 5
	testNum := advertDF.Nrow() / 5

	if traningNum+testNum < advertDF.Nrow() {
		traningNum++
	}

	// creating a subset indices.
	trainingIdx := make([]int, traningNum)
	testIdx := make([]int, testNum)

	// Enumerate the traning indices

	for i := 0; i < traningNum; i++ {
		trainingIdx[i] = i
	}

	// Enumerate test indices
	for i := 0; i < testNum; i++ {
		testIdx[i] = traningNum + i
	}

	// create a the subset dataframes.
	// Every subset will cut the data frame in portions
	// depend of the number of rows that you will set
	trainingDF := advertDF.Subset(trainingIdx)
	testDF := advertDF.Subset(testIdx)

	// Create a mapa that will be used in writing the data
	// to files

	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// create the respective files

	for idx, setName := range []string{"training.csv", "test.csv"} {

		// save the filtered dataset file.
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		// create a buffered writer
		w := bufio.NewWriter(f)

		// Write the daraframe out as a CSV

		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}

}

func StartTrainingModel(pathFile string) {

	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create a new CSV reader reading from the opened file

	reader := csv.NewReader(file)

	// Read in all of the CSV records

	reader.FieldsPerRecord = 4

	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// in this case we are going to try and model our sales(y)
	// by the TV feature plus an intercep. As such, let's create
	// the struct needed to train a model using Sajari repository

	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")

	// Now loop of records in the csv, adding the training
	// data to the refression value

	for i, record := range trainingData {

		// skip the header
		if i == 0 {
			continue
		}

		// parse the Sales regression measure, or "y"

		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// parse the Tv value.

		tvVal, err := strconv.ParseFloat(record[0], 64)

		if err != nil {
			log.Fatal(err)
		}

		// Add these points to the regression value.
		r.Train(regression.DataPoint(yVal, []float64{tvVal}))
	}

	// Train/fit the regression model.
	r.Run()

	// Output the trained model parameters.
	fmt.Printf("\nRegression formula: \n%v\n\n", r.Formula)

}

func main() {
	//pathFile := "Advertising.csv"
	pathTraining := "training.csv"
	//readingAdvertising(pathFile)
	//CheckingCorrelation(pathFile)
	//CreatingDataTraining(pathFile)
	StartTrainingModel(pathTraining)
}
