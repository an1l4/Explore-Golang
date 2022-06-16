package main

import "fmt"

func myFunction(x int,y string) (result int, text string) {
	result = x + x
	text = y + "World"
	return
}

func main() {
	fmt.Println(myFunction(10,"Hello"))
}