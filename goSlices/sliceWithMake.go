package main

import "fmt"

func slice(s string, x []int) {
	fmt.Printf("%s len = %d, cap = %d %v \n",s,len(x),cap(x),x)
}

func main() {
	a := make([]int,5)
	slice("a",a)

	b := make([]int,0,5)
	slice("b",b)

	c :=b[:2]
	slice("c",c)

	d := c[2:5]
	slice("d",d)
}