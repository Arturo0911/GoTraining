package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Arturo0911/GoTraining/HeadFirstGo/functions"
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
	data := [4]int{5, 3, 56, 8}
	if err != nil {
		panic(err)
	}

	for _, value := range data {
		fmt.Println(value)
	}

	for _, lines := range fileLines {
		fmt.Println(lines)
	}

	functions.Making()

	//fmt.Println(fileLines)
}
