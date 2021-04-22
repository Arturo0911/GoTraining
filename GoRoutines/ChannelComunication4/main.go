package main

import (
	"fmt"
	"sync"
)

func workers(ports chan int, wg *sync.WaitGroup) {

	for x := range ports {
		fmt.Println(x)
		wg.Done()
	}

}

func main() {

	var wg sync.WaitGroup
	ports := make(chan int, 100)

	for x := 0; x < cap(ports); x++ {

		go workers(ports, &wg)
	}

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
	close(ports)

}
