package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
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
* @return slices with the data inside
 */
func ReaderFile(fileName string) [][]string {

	recordFile, err := os.Open(fileName)

	if err != nil {

		fmt.Println(err)
	}

	reader := csv.NewReader(recordFile)

	records, _ := reader.ReadAll()

	//fmt.Println(records)
	fmt.Println(reflect.ValueOf(records).Kind())
	/*for _, value := range records {

		if value[4] == "Iris-setosa" {
			fmt.Println(value)
			break
		}

	}*/

	return records

}

func main() {
	ReaderFile("datasets/iris.csv")
}
