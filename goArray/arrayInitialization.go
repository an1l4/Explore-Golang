package main

import "fmt"

func main() {
	array1 := [5]int{}
	array2 := [5]int{1,2,3}
	array3 := [5]int{1,2,3,4,5}

	//if no elements it assign with default valu 0 for int, "" for string
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	//initialize only specific elements
	array4 := [5]int{1:20,2:40}
	fmt.Println(array4)
}