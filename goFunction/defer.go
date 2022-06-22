package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three") //LIFO last in first out

	fmt.Println("Hello")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
