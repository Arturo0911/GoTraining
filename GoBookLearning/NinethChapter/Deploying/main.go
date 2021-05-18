package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
)

// Declare the input and output directory flags

var (
	inDirPtr  = flag.String("inDir", "", "The directory containing the trainig data")
	outDirPts = flag.String("outDir", "", "The output directory")
)

const pathFile = "../Advertising.csv"
const trainFile = "train.csv"
const testFile = "test.csv"

// ModelInfo includes the information about the
// model that is putput from the training
type ModelInfo struct {
	Intercept    float64           `json:"intercept"`
	Coefficients []CoefficientInfo `json:"coefficients"`
}

// CoefficientsInfo include information about a
// particular model coefficient.
type CoefficientInfo struct {
	Name        string  `json:"name"`
	Coefficient float64 `json:"coefficient"`
}

func createDataTrainTest() {
	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dataDF := dataframe.ReadCSV(file)

	// Crating the num of the train and test data
	trainNum := (4 * dataDF.Nrow()) / 5
	testNum := dataDF.Nrow() / 5

	if trainNum+testNum < dataDF.Nrow() {
		trainNum++
	}

	trainSet := make([]int, trainNum)
	testSet := make([]int, testNum)

	for i := 0; i < trainNum; i++ {
		trainSet[i] = i
	}

	for j := 0; j < testNum; j++ {
		testSet[j] = j + trainNum
	}

	trainSubset := dataDF.Subset(trainSet)
	testSubset := dataDF.Subset(testSet)

	modelMap := map[int]dataframe.DataFrame{
		0: trainSubset,
		1: testSubset,
	}

	for idx, colName := range []string{"train.csv", "test.csv"} {

		f, err := os.Create(colName)
		if err != nil {
			log.Fatal(err)
		}

		w := bufio.NewWriter(f)

		if err := modelMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}

func trainTestModel() {
	newTrainFile, err := os.Open(trainFile)
	if err != nil {
		log.Fatal(err)
	}
	defer newTrainFile.Close()

	reader := csv.NewReader(newTrainFile)
	reader.FieldsPerRecord = 4

	dataReader, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")

	for i, record := range dataReader {

		// Skiping the index
		if i == 0 {
			continue
		}

		// Parsing the data
		yVals, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		tvVals, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		r.Train(regression.DataPoint(yVals, []float64{tvVals}))
	}

	// Training and fit the regression model
	r.Run()

	fmt.Printf("\nFormula from the predicted value %v\n\n", r.Formula)

	newTestFile, err := os.Open(testFile)
	if err != nil {
		log.Fatal(err)
	}
	defer newTestFile.Close()

	testReader := csv.NewReader(newTestFile)
	testReader.FieldsPerRecord = 4

	testData, err := testReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var mAE float64

	for i, colVal := range testData {

		// skiping the header
		if i == 0 {
			continue
		}

		yObserved, err := strconv.ParseFloat(colVal[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		tvValObserved, err := strconv.ParseFloat(colVal[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		yPredicted, err := r.Predict([]float64{tvValObserved})
		if err != nil {
			log.Fatal(err)
		}

		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}

	fmt.Printf("\nMAE = %0.3f\n\n", mAE)
}

func main() {
	flag.Parse()

	_, err1 := os.Open(trainFile)
	_, err2 := os.Open(testFile)
	if err1 != nil && err2 != nil {
		log.Printf("[*] Files Doesn't exists; creating files...")
		createDataTrainTest()
	}
	trainTestModel()

	// Train/ Fit the singled regression model
	// with Sajari regression

	// Fill in the model information.

}
