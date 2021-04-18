package main

import (
	"fmt"
)

type Behavior interface {
	Breath() bool // -> is Mamifer or not!
	Walk() bool
	Swim() bool
}

type Animal struct {
	Name   string
	Type   string
	Age    int
	Height float64
	Status string
}

func (p Animal) Breath() bool {

	if p.Type == "Fish" {
		return false
	} else {
		return true
	}

}

func Load(a *Animal) {

	a.Name = "Arturo"
	a.Height = 28
	a.Height = 181
	a.Status = "Single"
	a.Type = "Human"

}

func main() {
	var a Animal
	Load(&a)
	fmt.Println(a)
}
