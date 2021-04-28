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

	for x := len(arr) - 1; x >= 0; x-- {
		for y := len(arr) - 1; y > len(arr)-1-x; y-- {

			if arr[y] < arr[y-1] {
				//fmt.Println(arr[y], arr[y-1])
				arr[y], arr[y-1] = arr[y-1], arr[y]
			}
		}
	}

	if len(arr)%2 != 0 {

		// position number normal
		mediamPos = (len(arr) / 2) + 1
		value = arr[mediamPos-1]
		lowQuartile = (arr[((len(arr)-mediamPos)/2)] + arr[((len(arr)-mediamPos)/2)-1]) / 2
		highQuartile = (arr[(mediamPos+((len(arr)-mediamPos)/2))-1] + arr[(mediamPos+((len(arr)-mediamPos)/2))]) / 2

		quartile = append(quartile, lowQuartile)
		quartile = append(quartile, value)
		quartile = append(quartile, highQuartile)
		return quartile

	} else {
		mediamPos = (len(arr) / 2) - 1
		mediamPos2 := (len(arr) / 2)
		value = (arr[mediamPos] + arr[mediamPos2]) / 2
		lowQuartile = (arr[((len(arr)-mediamPos)/2)] + arr[((len(arr)-mediamPos)/2)-1]) / 2
		highQuartile = (arr[(mediamPos+(len(arr)-mediamPos2)/2)-1] + arr[(mediamPos2+(len(arr)-mediamPos)/2)]) / 2

		quartile = append(quartile, lowQuartile)
		quartile = append(quartile, value)
		quartile = append(quartile, highQuartile)

		return quartile
	}

}

func main() {

	//arr := []int32{3, 7, 8, 5, 12, 14, 21, 13, 18}
	arr := []int32{3, 7, 8, 5, 12, 14, 21, 15, 18, 14}
	fmt.Println(arr)
	fmt.Println(quartiles(arr))
	//fmt.Println(quartiles(arr2))
}
