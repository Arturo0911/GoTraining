package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	s := "11:59:59PM"
	finalHour := ""
	splited := strings.Split(s, "")

	//fmt.Println(splited)
	//fmt.Println(reflect.TypeOf(splited))

	//number, err := strconv.Atoi(splited[1])

	/*if err != nil {
		fmt.Println(err)
	}

	fmt.Println(12 + number)*/

	for x := 1; x < len(splited)-2; x++ {

		if splited[len(splited)-2] == "A" && splited[0] != "1" {

			finalHour = splited[0] + splited[1] + splited[2] + splited[3] + splited[4] + splited[5] + splited[6] + splited[7]
			break

		} else {
			if x == 1 {

				if splited[x-1] == "0" {

					number, err := strconv.Atoi(splited[x])

					if err != nil {
						fmt.Println(err)
					}

					finalHour += strconv.Itoa(number + 12)

				} else {

					finalHour += "00"
				}

			} else {

				finalHour += splited[x]

			}

			//finalHour += splited[x]
		}
	}

	fmt.Println(finalHour)

}
