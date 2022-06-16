package main

import "fmt"

func main() {
	//variable declaration with initial value using var keyword and mentioning the type
	var student1 string = "John"

	//variable decalration with initial value using var keyword and without mentioning type
	var student2 = "Jane" //here datatype will taken from the value "Jane" so its a string

	//variable declarion with initial value using ":=" sign
	x := 2 //here no need to mention the var keyword and datatype, ":=" sign need to assign the value as well,only can be used inside the function

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}