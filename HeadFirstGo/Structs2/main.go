package main


import "fmt"


type suscriber struct {

	name 	string
	rate 	float64
	active 	bool
}

func applyDiscount(s *suscriber){

	s.rate = 4.99
}



func main(){


	fmt.Println("hi!!")
	var s suscriber

	applyDiscount(&s)

	fmt.Println(s.rate)

	var value int = 2

	var pointer *int = &value

	fmt.Println(pointer)
	fmt.Println(*pointer)
}