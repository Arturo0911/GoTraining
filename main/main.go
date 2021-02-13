package main

/*
* Started with Structs
 */

type Product struct {
	id           int
	productName  string
	productStock int
	productPrice float64
}

type Country struct {
	name     string
	quantity int
}

func main() {

	var country1, country2, country3 Country

	country1.name = "Ecuador"
	country1.quantity = 18000000

	country2.name = "Peru"
	country2.quantity = 32000000

	country3.name = "Brazil"
	country3.quantity = 50000000
	

	if country1.quantity > country2.quantity && country2.quantity > country3.quantity{

		fmt.Println(country1.name, " is the country with the most people inside"
)

	}else {


		if country2.quantity



	}

 



}
