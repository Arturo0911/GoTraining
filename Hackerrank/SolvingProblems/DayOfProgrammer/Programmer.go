package main

import (
	"fmt"
)

/**
* @author Pxyl0xd
* @description Hackerrank solutions
 */

func IsBisiest(year int32) bool {

	response := true
	if year%4 == 0 && year%100 != 0 {
		response = true
	} else {
		response = false
	}

	return response

}

func LoadDays(is bool) map[string]int {
	hashYear := make(map[string]int)
	hashYear["January"] = 31
	hashYear["February"] = 28
	hashYear["March"] = 31
	hashYear["April"] = 30
	hashYear["May"] = 31
	hashYear["June"] = 30
	hashYear["July"] = 31
	hashYear["August"] = 31
	hashYear["September"] = 30
	hashYear["October"] = 31
	hashYear["November"] = 30
	hashYear["December"] = 31

	if is {
		return hashYear
	} else {
		hashYear["February"] += 1
		return hashYear
	}

}

func dayOfProgrammer(year int32) string {

	if year == 1918 {
		return "26.09.1918"
	} else if (year <= 1917 && year%4 == 0) || (year%4 == 0 && year%400 == 0) || (year%4 == 0 && year%100 != 0) {

		return fmt.Sprintf("12.09.%v", year)

	} else {

		return fmt.Sprintf("13.09.%v", year)
	}

}

func Testing() {

	year1 := 2016
	year2 := 2021
	year3 := 2008
	year4 := 2017
	year5 := 2000
	year6 := 2500
	year7 := 1988
	year8 := 1992
	year9 := 1800

	fmt.Println(year1, dayOfProgrammer(int32(year1)))
	fmt.Println(year2, dayOfProgrammer(int32(year2)))
	fmt.Println(year3, dayOfProgrammer(int32(year3)))
	fmt.Println(year4, dayOfProgrammer(int32(year4)))
	fmt.Println(year5, dayOfProgrammer(int32(year5)))
	fmt.Println(year6, dayOfProgrammer(int32(year6)))
	fmt.Println(year7, dayOfProgrammer(int32(year7)))
	fmt.Println(year8, dayOfProgrammer(int32(year8)))
	fmt.Println(year9, dayOfProgrammer(int32(year9)))

}

func main() {
	Testing()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
