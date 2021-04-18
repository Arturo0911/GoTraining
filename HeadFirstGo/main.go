package main

import (
	"bufio"
	"fmt"
	"os"
)

func getStrings(fileName string) ([]string, error) {

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

	return lines, nil

}

func main() {

	fileLines, err := getStrings("golang.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(fileLines)
}
