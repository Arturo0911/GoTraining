package main

import (
	"fmt"
	"time"
)

func loop(channel chan int) {

	for x := 0; x < 100000000; x++ {

		time.Sleep(1 * time.Millisecond)
		fmt.Println(x)
		//channel <- x
	}

	channel <- 10

}

// Got it
func main() {

	myChannel := make(chan int)

	go loop(myChannel)
	//fmt.Println(<-myChannel)
	go loop(myChannel)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)

}
