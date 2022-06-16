package main

import "fmt"

func main() {
	mySlice := []string{"Anila","Soman"}
	fmt.Println(mySlice)

	mySlice = append(mySlice,"John","Jane")
	fmt.Println(mySlice)
}