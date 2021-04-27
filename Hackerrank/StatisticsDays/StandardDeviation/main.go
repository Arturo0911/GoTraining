package main

import (
	"fmt"
	"math"
)

func stdDev(arr []int32) {
	// Print your answers to 1 decimal place within this function
	u := 0.0
	stdr := 0.0

	for _, value := range arr {

		u += float64(value)

	}

	mean := u / float64(len(arr))

	for _, element := range arr {

		stdr += float64(math.Pow(float64(element)-mean, 2))

	}

	fmt.Printf("%0.1f\n", math.Sqrt((stdr / float64(len(arr)))))

}

func main() {

	arr := []int32{10, 40, 30, 50, 20}
	stdDev(arr)
}
