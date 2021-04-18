package main

/*
* @author Arturo Negreiros
* @description Maps and Structs
 */

import (
	"fmt"
)

func main() {

	ranks := make(map[string]int)
	ranks["Arturo's age"] = 28
	ranks["Sarah's age"] = 3

	for key, value := range ranks {

		fmt.Println(key, value)
	}

	fmt.Println(ranks)
}
