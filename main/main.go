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

	/*var country1, country2, country3 Country

	country1.name = "Ecuador"
	country1.quantity = 18000000000000

	country2.name = "Peru"
	country2.quantity = 32000000000

	country3.name = "Brazil"
	country3.quantity = 50000000

	if country1.quantity > country2.quantity && country2.quantity > country3.quantity {

		fmt.Println(country1.name, " is the country with the most people inside")

	} else {

		if country2.quantity > country3.quantity && country3.quantity > country1.quantity {

			fmt.Println(country2.name, " is the country with the biggest quantity of people")

		} else if country3.quantity > country2.quantity && country2.quantity > country1.quantity {

			fmt.Println(country3.name, " is the country with the biggest quantity of people")
		}

	}*/

}
