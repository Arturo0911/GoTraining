package main

import "fmt"

/*
* Started with Structs
 */

type Product struct {
	id           int
	productName  string
	productStock int
	productPrice float64
}

type Country struct {
	name     string
	quantity int
}

type Person struct {
	credential  string
	name        string
	lastName    string
	phonenumber string
	age         int
}

func loadData(person *Person) {

	person.credential = "0918237421"
	person.name = "Arturo Francesco"
	person.lastName = "Negreiros Samanez"
	person.phonenumber = "0963951028"
	person.age = 29

}

func printPerson(person Person) {

	fmt.Println(person)
}

func main() {

	var person Person

	loadData(&person)
	printPerson(person)
}
