package main

import "fmt"

func main() {
	mySlice1 := []int{}
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
	fmt.Println(mySlice1)

	mySlice2 := []string{"Go","Slices","are","Powerful"}
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
	fmt.Println(mySlice2)
}