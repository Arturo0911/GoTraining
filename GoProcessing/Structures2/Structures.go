package main

import "fmt"

type Square struct {
	Height float64
}

type MathProcess interface {
	AreaCalc() float64
	PerimCalc() float64
}

func (s Square) AreaCalc() float64 {
	return (s.Height * s.Height)
}

func (s Square) PerimCalc() float64 {

	return (s.Height * 4)
}

func main() {

	square := Square{
		Height: 15,
	}
	ShowingInformation(square)
}

func ShowingInformation(m MathProcess) {

	fmt.Println(m.AreaCalc())
	fmt.Println(m.PerimCalc())

}
