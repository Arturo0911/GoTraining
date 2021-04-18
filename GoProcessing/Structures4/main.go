package main


import "fmt"

func countChannel(firstChan chan int){
	x:= 0
	for {
		firstChan <- x
		x++
	}
}

func Printing(secondChan chan int){
	var value int
	for {
		value = <- secondChan
		fmt.Println(value)
	}
}

/**
 * making Goroutines and 
 * comunication between channels
 */

func main(){

	channel := make(chan int)
	go countChannel(channel)
	go Printing(channel)
	var final string
	fmt.Scanln(&final)
}