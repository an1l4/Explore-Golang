package main

import "fmt"

func slice(s []int) {
	fmt.Printf("len = %d, capacity = %d %v \n",len(s),cap(s),s)
}

func main() {
	s := []int{2,3,5,7,11,13}
	slice(s)

	//zero length 
	s = s[:0]
	slice(s)

	//extend length
	s = s[:4]
	slice(s)

	//drop first two values
	s = s[2:]
	slice(s)
}