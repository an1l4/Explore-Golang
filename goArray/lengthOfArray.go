package main

import "fmt"

func main() {
	cars := [4]string{"Polo","i20","Figo","XUV700"}
	numbers := [...]int{1,2,3}

	fmt.Println(len(cars))
	fmt.Println(len(numbers))
}