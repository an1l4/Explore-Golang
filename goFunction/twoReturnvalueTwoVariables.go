package main

import "fmt"

func myFunction(x int, y string) (result int, text string) {
	result = x + x
	text = y + "world"
	return
}

func main() {
	a,b := myFunction(5,"hello")
	fmt.Println(a,b)
}