package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var name string
	in := bufio.NewReader(os.Stdin)
	name = in.ReadString()
	fmt.Println(name)
	//var names string
	//var address string

	//fmt.Println(names, " vive en: ", address)
}
