package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// Matrix Samples
func MakingMatrix() {
	data := []float64{1.2,5.7,2.4,7.3}
	
	a := mat.NewDense(2,2,data)
	fa := mat.Formatted(a, mat.Prefix(" "))
	fmt.Printf("mat = %v\n\n", fa)
	

	// we can access and modify certain values withing A via build-in methods:
	// access to position begint for zero as normal index in Informatic.

	val := a.At(0, 1)
	fmt.Println(val)
}

func main() {
	MakingMatrix()
}
