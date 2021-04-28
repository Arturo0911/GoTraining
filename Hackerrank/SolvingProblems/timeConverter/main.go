package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	finalHour := ""
	splited := strings.Split(s, "")
	fmt.Println(s)

	for x := 1; x < len(splited)-2; x++ {

		if splited[8]+splited[9] == "PM" {

			if splited[0]+splited[1] == "12" {
				finalHour = splited[0] + splited[1] + splited[2] + splited[3] + splited[4] + splited[5] + splited[6] + splited[7]
				break
			} else {
				if x == 1 {

					if splited[x-1]+splited[x] != "12" {

						number, err := strconv.Atoi(splited[x-1] + splited[x])

						if err != nil {
							fmt.Println(err)
						}

						finalHour += strconv.Itoa(number + 12)

					}

				} else {

					finalHour += splited[x]

				}
			}

		} else if splited[8]+splited[9] == "AM" {

			if splited[x-1]+splited[x] == "12" {
				finalHour = "00" + splited[2] + splited[3] + splited[4] + splited[5] + splited[6] + splited[7]
				break
			} else {
				finalHour = splited[0] + splited[1] + splited[2] + splited[3] + splited[4] + splited[5] + splited[6] + splited[7]
				break
			}
		}

	}

	return finalHour

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
