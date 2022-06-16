package main

import "fmt"

type Person struct {
	name string
	age int
	salary int

}

func main() {
	var per1 Person
	var per2 Person
	
	per1.name = "John"
	per1.age = 20
	per1.salary = 2000

	per2.name = "Jane"
	per2.age = 30
	per2.salary = 4000

	printPerson(per1)
	printPerson(per2)

}
func printPerson(per Person) {
	fmt.Println(per.name)
	fmt.Println(per.age)
	fmt.Println(per.salary)
}