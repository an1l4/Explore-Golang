package main

import "fmt"

func main() {
	mySlice := make([]int,5,10)
	fmt.Printf("mySlice = %v \n",mySlice)
	fmt.Printf("length = %d \n",len(mySlice))
	fmt.Printf("capacity = %d \n",cap(mySlice))

	//omitted capacity
	mySlice1 := make([]int,5)
	fmt.Printf("mySlice1 = %v \n",mySlice1)
	fmt.Printf("length = %v \n",len(mySlice1))
	fmt.Printf("capacity = %v \n",cap(mySlice1))
}