package main

import "fmt"

func main() {
	i,j := 42,2701

	p := &i  //point to p
	fmt.Println(*p) //read i through the pointer
	*p = 21 // set i through pointer
	fmt.Println(i) //see new value of i

	p = &j //point to p
	*p = *p / 37 //divide j through pointer
	fmt.Println(j) //see value of j
}