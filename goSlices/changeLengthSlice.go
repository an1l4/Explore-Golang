package main

import "fmt"

func main() {
	array := [6]int{1,2,3,4,5,6}
	fmt.Println(array)

	mySlice := array[1:5]
	fmt.Println(mySlice)

	mySlice = array[1:3]
	fmt.Println(mySlice)

	mySlice = append(mySlice,10,20,30,40)
	fmt.Println(mySlice)

}