package main

import (
	"fmt"
	"strings"
	"time"
)

// source of information
// programacion Go ya
// https://www.tutorialesprogramacionya.com/goya/
// https://www.udemy.com/course/lenguaje-go

func myNameSlow(name string) {

	letras := strings.Split(name, "")
	//fmt.Println(letras)

	for _, letra := range letras {
		time.Sleep(1 * time.Second)
		fmt.Println(letra)
	}

}

// func loop(channel1 chan time.Duration) {
// 	init := time.now()

// 	for i := 0; i < 1000000000000; i++ {

// 	}

// 	final := time.Now()

// 	channel1 <- final.Sub(init)

// }

func main() {

	myNameSlow("Arturo")

	channel1 := make(chan time.Duration)

	go loop(channel1)
	fmt.Println("i'm here")

}
