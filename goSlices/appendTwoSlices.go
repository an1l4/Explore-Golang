package main

import "fmt"

func main() {
	mySlice1 := []int{1,2,3}
	mySlice2 := []int{4,5,6}
	mySlice3 := append(mySlice1, mySlice2...)
	//The '...' after slice2 is necessary when appending the elements of one slice to another.

	fmt.Println(mySlice3)
}