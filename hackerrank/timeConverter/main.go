package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	name := "07:05:45PM"
	splited := strings.Split(name, "")

	fmt.Println(splited)
	fmt.Println(reflect.TypeOf(splited))

	fmt.Println(12 + int(splited[1]))
}
