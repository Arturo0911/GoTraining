package main

import (
	"fmt"
)

func quartiles(arr []int32) []int32 {
	// Write your code here

	var quartile []int32
	var highQuartile int32
	var lowQuartile int32

	var mediamPos int
	var value int32 = 0

	for i := len(arr); i >= 0; i-- {

		for j := 0; j > len(arr)-i; j-- {

			if arr[j] > arr[j-1] {

				arr[j-1], arr[j] = arr[j], arr[j-1]

			}

		}
	}

	fmt.Println(arr)

	if len(arr)%2 != 0 {

		// position number normal
		mediamPos = (len(arr) / 2) + 1
		value = arr[mediamPos-1]
		lowQuartile = (arr[(mediamPos/2)-1] + arr[(mediamPos/2)]) / 2
		highQuartile = (arr[(mediamPos+(len(arr)-mediamPos)/2)-1] + arr[(mediamPos+(len(arr)-mediamPos)/2)]) / 2

		quartile = append(quartile, lowQuartile)
		quartile = append(quartile, value)
		quartile = append(quartile, highQuartile)
		return quartile

	} else {
		mediamPos = (len(arr) / 2)
		mediamPos2 := (len(arr) / 2) + 1
		lowQuartile = (arr[(mediamPos/2)-1] + arr[(mediamPos/2)]) / 2
		highQuartile = (arr[(mediamPos+(len(arr)-mediamPos2)/2)-1] + arr[(mediamPos2+(len(arr)-mediamPos)/2)]) / 2

		quartile = append(quartile, lowQuartile)
		quartile = append(quartile, value)
		quartile = append(quartile, highQuartile)

		return quartile
	}

}

func main() {

	arr := []int32{3, 7, 8, 5, 12, 14, 21, 13, 18}
	//arr2 := []int32{3, 7, 8, 5, 12, 14, 21, 13, 18, 20}
	fmt.Println(quartiles(arr))
	//fmt.Println(quartiles(arr2))
}
