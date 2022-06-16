package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(2,3))
	fmt.Println(sum(2,2))
}