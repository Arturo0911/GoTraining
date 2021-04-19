package main


import "fmt"


type suscriber struct {

	name 	string
	rate 	float64
	active 	bool
}



func printInformation(s *suscriber){

	fmt.Println("name: ", s.name)
	fmt.Println("rate: ", s.rate)
	fmt.Println("active: ", s.active)
}


func loadSuscriber(name string) *suscriber{

	var s suscriber
	s.name = name
	s.rate = 5.99
	s.active = true

	return &s

}



func applyDiscount(s *suscriber){

	s.rate = 4.99
}

func banner(){
	fmt.Println("-------------------------------\n")
}


func main(){
	banner()
	suscriber1 := loadSuscriber("Arturo Negreiros")
	applyDiscount(suscriber1)

	printInformation(suscriber1)
	banner()
	suscriber2 := loadSuscriber("Leslie Chavez")

	printInformation(suscriber2)

	banner()
}