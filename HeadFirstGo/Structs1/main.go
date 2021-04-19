package main

/**
 *@author Arturo Negreriros
 */


import "fmt"


/*Struct sections*/
type Person struct {

	Name string
	Age int
	Email string
}


type myStruct struct {

	number float64
	name string
	toggle bool
}

type part struct {
	description string
	count int
}

type car struct {
	name string
	topSpeed float64
}

func showInfoPart(p part){

	fmt.Println(p.description)
	fmt.Println(p.count)
}


func banner (){

	fmt.Println("------------------------------\n")
}

/*Creating the self function to manage the creating of
	type and her assignation
*/


func cratePart(description string, counter int) part{

	var p part

	p.description = description
	p.count = counter

	return p
}
/*
* Modifying the value from the struct type, this one
* should be shiped with the pointer to access the data
* and modifying it, otherwise, you only gonna modify the copy
 */

//

type suscribers struct {
	name string
	rate float64
}


func changeStruct(s suscribers) {
	s.name = "Arturo Negreiros"

}



func loadData() Person{

	persona := Person {
		Name: "Arturo Negreiros",
		Age: 28,
		Email: "anegreiross@outlook.com",
	}	

	return persona
}


func main(){


	var s suscribers
	changeStruct(s)
	fmt.Println(s)




	banner()

	fmt.Println("Hii")
	fmt.Println(loadData())

	var myStructure myStruct

	myStructure.number = 65

	fmt.Println(myStructure.number)

	banner()


	var chevrolet car

	chevrolet.name = "Grand Vitara"
	chevrolet.topSpeed = 220

	fmt.Println("Car name: ", chevrolet.name)
	fmt.Println("Top speed: ", chevrolet.topSpeed)
	banner()



	var p part

	p.description = "car part"
	p.count	 = 56
	showInfoPart(p)

	banner()


	showInfoPart(cratePart("Clutch box", 50))
	banner()
	showInfoPart(cratePart("Carrousel", 255))
	banner()
	showInfoPart(cratePart("Automatic windows", 300))

	banner()
	showInfoPart(cratePart("scape pipe", 175))
}