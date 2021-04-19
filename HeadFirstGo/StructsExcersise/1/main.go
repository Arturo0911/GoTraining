package main

import "fmt"

// structs

type car struct {
	name     string
	topSpeed float64
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func main() {
	var mustang car
	mustang.name = "Mustang Cobra"
	mustang.topSpeed = 225

	// the '&' was added because nitroBoost function
	// didn't modify the value, for this reason
	// pass a pointer
	nitroBoost(&mustang)

	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)
}
