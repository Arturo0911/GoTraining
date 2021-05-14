package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const pathFile = "../datasets/fleet_data.csv"

func Histogram() {
	clusFile, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
	}
	defer clusFile.Close()

	driverDF := dataframe.ReadCSV(clusFile)

	fmt.Println(driverDF.Describe())

	for _, colName := range driverDF.Names() {

		if colName == "Driver_ID" {
			continue
		}

		// First step
		// making the plot vals taking as reference the number of rows
		plotVals := make(plotter.Values, driverDF.Nrow())
		for i, plotVal := range driverDF.Col(colName).Float() {
			plotVals[i] = plotVal
		}

		// Second step
		// Creating a new plot to be shiped
		p := plot.New() // This one only take one response, another old versions, this take an error parameter
		p.Title.Text = fmt.Sprintf("Histogram of %s", colName)

		// third step
		// makes the kinna plot that you wanna create
		pHist, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}
		pHist.Normalize(1) // Normalize makes reference to the data in normal order

		// Four step
		// add the histogram to the plot
		p.Add(pHist)

		// Five step
		// save the values in png files

		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hitogram.png"); err != nil {
			log.Fatal(err)
		}
	}
}

func NewPlotting() {

	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatal("error file: ", err)
	}
	defer file.Close()

	driverDF := dataframe.ReadCSV(file)

	// Distance
	yVals := driverDF.Col("Distance_Feature").Float()

	pts := make(plotter.XYs, driverDF.Nrow())

	// Fill pts with data
	for i, floatVal := range driverDF.Col("Speeding_Feature").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
	}

	p := plot.New()

	p.X.Label.Text = "Speeding"
	p.Y.Label.Text = "Distance"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	// Added color and style
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	// Saving the point in a PNG file
	p.Add(s)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "fleet_data_scatter.png"); err != nil {
		log.Fatal(err)
	}

}

func main() {

	//Histogram()
	NewPlotting()
}
