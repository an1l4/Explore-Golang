package main

import "fmt"

func main() {
	myMessage := make(chan int, 3)

	myMessage <- 1
	myMessage <- 2
	myMessage <- 3

	fmt.Println(<-myMessage)
	fmt.Println(<-myMessage)
	fmt.Println(<-myMessage)
}