package main

import "fmt"


/*Interfaces*/

type Behavior interface {
	
	fly()

}

type Airplane struct {

	name string
}

type Chooper struct {
	
	name string

}


func (a Airplane) fly(){
	fmt.Println("the  ", a.name, "is flying")
}

func (c *Chooper) fly(){
	
	fmt.Println("the ", c.name, " is flying")

}












/*Function with two values to return*/

func HigherLower(value1, value2 int)(int, int){
	
	
	if value1 >= value2{
		return value1, value2
	}else {
		return value2, value1
	}


}	



func main() {
	
	a := Airplane {
		
		name: "Airplane",
	
	}

	a.fly()

	//fmt.Println(HigherLower(25, 35))
	
}
