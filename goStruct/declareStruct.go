package main

import "fmt"

type Person struct {
	name string
	age int
	job string
	salary int
}

func main() {
	var per1 Person
	var per2 Person

	per1.name = "John"
	per1.age = 20
	per1.job = "teacher"
	per1.salary = 20000

	per2.name = "Jane"
	per2.age = 30
	per2.job = "marketing"
	per2.salary = 30000

	//accessing struct members
	fmt.Println("Name :",per1.name)
	fmt.Println("Age :",per1.age)
	fmt.Println("job :",per1.job)
	fmt.Println("salary :",per1.salary)

	fmt.Println("Name :",per2.name)
	fmt.Println("Age :",per2.age)
	fmt.Println("job :",per2.job)
	fmt.Println("salary :",per2.salary)
}