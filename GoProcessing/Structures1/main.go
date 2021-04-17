package main

import "fmt"

/*Interfaces*/

type Behavior interface {
	Fly() string
}

type Airplane struct {
	name string
}

type Chooper struct {
	name string
}

func (a Airplane) Fly() string {
	return "the  " + a.name + "is flying"
}

func (c Chooper) Fly() string {

	return fmt.Sprintf("The %s is flying", c.name)
}

/*func (c *Chooper) fly() string {

	return "the " + c.name + " is flying"

}*/

func main() {

	a := Airplane{
		name: "Airplane",
	}

	c := Chooper{
		name: "Chooper",
	}

	Print(a)
	Print(c)

}

func Print(b Behavior) {
	fmt.Println(b.Fly())
}
