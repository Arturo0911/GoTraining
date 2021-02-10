package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	var valor int
	suma := 0

	for {

		fmt.Println("Input 0 to finish the loop")
		fmt.Scan(&valor)

		if valor == 0 {
			break
		}

		suma += valor

	}

	fmt.Println("El valor de la suma de todos los números es: ", suma)

}
