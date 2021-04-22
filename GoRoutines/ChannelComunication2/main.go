package main

import (

	"fmt"
	"strings"
	"time"
)



// source of information
// programacion Go ya
// https://www.tutorialesprogramacionya.com/goya/


func myNameSlow(name string){
	
	letras := strings.Split(name, "")
	//fmt.Println(letras)

	for _, letra := range letras {
		time.Sleep(1 *time.Second)
		fmt.Println(letra)
	}


}





func main() {
	
	myNameSlow("Arturo")
}
