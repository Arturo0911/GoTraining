package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {

	e := Employee{}
	fmt.Printf("%v\n", e)
	//e.id = 1
	e.name = "Arturo Negreiros"
	fmt.Printf("%v\n", e)

	e.SetId(5)
	e.SetName("Payload")
	fmt.Printf("%v\n", e)

	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

	e2 := Employee{
		id:       1,
		name:     "arturon",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)

	e3 := new(Employee)

	fmt.Printf("%v\n", *e3)
	e3.id = 1
	e3.name = "Arturon"
	e3.vacation = true

	fmt.Printf("%v\n", e3)

	e4 := NewEmployee(27, "payload", true)
	fmt.Printf("%v\n", *e4)

}
