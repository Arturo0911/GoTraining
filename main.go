package main

import (
	"fmt"
	"reflect"
)

/*
* Specify what kind of value will be return the function
* in this case is type Integer
 */
func sumNumbers(number int) (int, string) {

	return number * 2, "Arturon"
}

func showNames(name string) string {

	return "Hello " + name + " it's a pleasure to meet you"
}

func declareArray(size int) []int {

	slice := make([]int, size)

	for i := 0; i < size; i++ {

		slice[i] = i
	}

	return slice

}

func findPrimeNumber(number int) bool {

	counter := 0
	for x := 2; x < number; x++ {

		if number%x == 0 {

			counter++
		}
	}

	if counter > 0 {

		return false

	} else {

		return true
	}

}

func createArray() [10]int {

	var array [10]int
	for x := 0; x < 10; x++ {

		array[x] = x
	}

	return array

}

func updateArray(array [10]int, newValue int, position int) [10]int {

	if position >= 10 {

		return array
	}

	array[position] = newValue

	return array

}

func searchPrimeNumbers(rangeNumber int) []int {

	primeList := make([]int, rangeNumber)
	position := 0
	for i := 2; i < rangeNumber; i++ {
		counter := 0
		for j := 2; j < rangeNumber; j++ {

			if i%j == 0 {

				counter++
			}

		}

		if counter < 2 {

			primeList[position] = i
			position++

		}

	}

	return primeList

}

func main() {

	// maps

	//countries := make(map[string]int)

	fmt.Println(searchPrimeNumbers(1000))

	var skills = make(map[string]string)

	skills["Java"] = "Backend"
	skills["Python"] = "Machine learning"
	skills["GraphQL"] = "Data transporting"
	skills["Arduino"] = "Robotic"

	var agesPerPerson = make(map[int]string)

	agesPerPerson[28] = "Arturon"
	agesPerPerson[43] = "Ines"
	agesPerPerson[41] = "Favio"

	fmt.Println(agesPerPerson)

	for a, b := range agesPerPerson {

		fmt.Println("key: ", a, " value: ", b)
	}

	/*countries["Ecuador"] = 1
	countries["Peru"] = 2

	fmt.Println(countries)
	fmt.Println(len(countries))*/

	fmt.Println(reflect.TypeOf(skills))

	// declare a slice

	quantity := 5
	bills := make([]int, quantity)

	for i := 0; i < 5; i++ {

		bills[i] = i
	}

	fmt.Println(bills)
	fmt.Println(reflect.TypeOf(bills))

	var array [5]int

	for x := 0; x < 5; x++ {

		array[x] = x
	}

	fmt.Println(array)
	fmt.Println(reflect.TypeOf(array))

	fmt.Println(declareArray(25))
	fmt.Println(reflect.TypeOf(declareArray(25)))

}

/*func main() {



	value, name := sumNumbers(25)

	fmt.Println(name)
	fmt.Println(value)

	names := showNames(name)

	fmt.Println(names)
	fmt.Println(findPrimeNumber(25))

	if findPrimeNumber(17) {
		fmt.Println("The number is Prime")

	} else {

		fmt.Println("The number IS NOT PRIME")

	}

	array := createArray()

	fmt.Println(array)
	fmt.Println(updateArray(array, 25, 8))

	fmt.Println("hello go")

	// Exist a lot of ways to built an array(data structured)

	var space [256]int

	for x := 0; x < 256; x++ {

		space[x] = x
	}

	fmt.Println(space)

	slice := space[100:150]

	fmt.Println(slice)

}*/
