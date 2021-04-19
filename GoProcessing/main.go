package main

import (
	"bufio"
	"fmt"
	"os"
)

// @author Arturo0911
// Trying to get datasets and use
// the value like data training

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

func main() {
	fmt.Println("hi everyone")
}
