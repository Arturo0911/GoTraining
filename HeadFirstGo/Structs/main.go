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






func loadData() Person{

	persona := Person {
		Name: "Arturo Negreiros",
		Age: 28,
		Email: "anegreiross@outlook.com",
	}	

	return persona
}


func main(){

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