package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {

	for i := 0; i < delay; i++ {

		fmt.Println(name, " sleeping")
		time.Sleep(1 * time.Second)

	}

}

func send(myChannel chan string) {

	reportNap("Sending Goroutine", 2)
	fmt.Println("***Sending value***")
	myChannel <- "a"
	fmt.Println("***Sending value***")
	myChannel <- "b"

}

func main() {

	myChannel := make(chan string)
	go send(myChannel)
	reportNap("receiving Goroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
