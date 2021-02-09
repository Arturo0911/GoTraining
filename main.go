package main

import "fmt"

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

func createArray(size int) [10]int {

	var array [10]int
	for x := 0; x < 10; x++ {

		array[x] = x
	}

	return array

}

func main() {

	/*var name string
	fmt.Scan(&name)
	*/

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

	fmt.Println(createArray(5))

	fmt.Println("hello go")
}
