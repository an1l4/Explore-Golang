package main

import "fmt"

func main() {
	//array decalration using var keyword
	var array1 = [3]int{1,2,3}
	
	//string
	var cars = [4]string{"Volvo","BMW","Ford","i20"}

	//array declaration using ":=" sign
	array2 := [5]int{4,5,6,7,8}

	//array have fixed length if length not mentioned it will calculate from number of values
	var array3 = [...]int{9,10,11}
	array4 := [...]string{"Anila","Soman"}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(cars)
}