package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math"
	"os"
)

// @author Arturo0911
// Trying to get datasets and use
// the value like data training

type ModelSet struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
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

/**
* @description Math process
* 				Making a Average, medium, moda, Variance
*
*
*		s²x = var(x)= 	∑n i=1	(xi−x^)²
*						---------------
*							n−1
*
 */

// Variance, standart desviation, average
func Variance(parameterList []float64, y []float64) (float64, float64, float64) {

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

func main() {
	fmt.Println(ReaderFile("datasets/iris.csv"))
}
