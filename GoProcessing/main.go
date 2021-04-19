package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

// @author Arturo0911
// Trying to get datasets and use
// the value like data training

type ModelSet struct {
	Variance           float64
	StandartDesviation float64
	PetalWidth         float64
	Species            string
}

func LoadDataSets(fileName string) ([]string, error) {
	//  in this case we gonna build a dataset with the file name or
	// to be more exactly, the path file
	var lines []string

	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()

	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, err
}

//
//	Variance formula
//  var(x)= ∑n i=1 (xi−x^)²
//			---------------
//				n−1

// Variance, standart desviation, average
func StatisticsValues(parameterList []float64) (float64, float64, float64) {

	var initSum float64 = 0
	var finalSum float64 = 0

	var newVariance float64 = 0
	var standartDesviation float64 = 0
	var average float64 = 0

	for _, value := range parameterList {
		initSum += value
	}

	average = (initSum / float64(len(parameterList)))

	for i := 0; i < len(parameterList); i++ {
		finalSum += math.Pow((parameterList[i] - average), 2)
	}

	newVariance = (finalSum / float64((len(parameterList) - 1)))
	standartDesviation = math.Pow(newVariance, 0.5)

	return average, newVariance, standartDesviation

}

func MakingXArrayValues(dataset [][]string, position int) ([]float64, error) {

	var parameterList []float64

	for _, value := range dataset {

		parameter, err := strconv.ParseFloat(value[position], 3)

		if err != nil {
			err := errors.New("cannot be parsing to float, because is not numeric parameter to parse to float")
			return nil, err
		}
		parameterList = append(parameterList, parameter)

	}

	return parameterList, nil
}

/**
* @return slices with the data inside
 */
func ReaderFile(fileName string) [][]string {
	recordFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(recordFile)
	records, _ := reader.ReadAll()

	return records

}

func weightedMean(X [5]int32, W [5]int32) {
	// Write your code here

	var sumTot int32 = 0
	var weightSum int32 = 0
	for i := 0; i < len(X); i++ {
		sumTot += (X[i] * W[i])
		weightSum += W[i]
	}

	fmt.Printf("%.1f\n", float32(sumTot/weightSum))

}

func main() {
	//fmt.Println(ReaderFile("datasets/iris.csv"))

	parameters, err := MakingXArrayValues(ReaderFile("datasets/iris.csv"), 0)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(StatisticsValues(parameters))

	x := [5]int32{10, 40, 30, 50, 20}
	w := [5]int32{1, 2, 3, 4, 5}

	weightedMean(x, w)
}
