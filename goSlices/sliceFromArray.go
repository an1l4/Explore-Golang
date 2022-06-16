package main

import "fmt"

func main() {
	myArray := [6]int{10,11,12,13,14,15}
	mySlice :=myArray[2:4]

	fmt.Printf("mySlice = %v \n",mySlice)
	fmt.Printf("Length = %d \n",len(mySlice))
	fmt.Printf("Capacity = %d \n",cap(mySlice))
}