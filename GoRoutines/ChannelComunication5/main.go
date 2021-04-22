package main


import (

	"fmt"
//	"time"

)



func print(c chan int){
	
	c <- 10

}


func anotherPrint(c chan int){

	
	sum := 0

	for i:=0; i < 100; i ++ {
	
		sum += i

	}

	c <- sum

}

func main(){

	c := make(chan int)
	
	go print(c)

	go anotherPrint(c)
	fmt.Println(<-c)

	fmt.Println(<-c)

}
