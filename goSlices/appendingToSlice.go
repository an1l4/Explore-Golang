package main

import "fmt"

func slice(s []int) {
	fmt.Printf("len = %d,cap = %d,%v \n",len(s),cap(s),s)
}

func main() {
	var s []int
	slice(s)

	//append works on nil slice
	s = append(s,0)
	slice(s)

	s = append(s,1)
	slice(s)

	s = append(s,2,3,4)
	slice(s)
}