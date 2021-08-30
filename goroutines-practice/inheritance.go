package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "This is a full time employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	TaxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "This is a temporary employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.age = 28
	ftEmployee.name = "payload"
	ftEmployee.id = 9
	fmt.Printf("%v\n", ftEmployee)

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
