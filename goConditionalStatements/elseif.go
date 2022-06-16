package main

import "fmt"

func main() {
	time := 22
	if time < 18 {
		fmt.Println("Good Morning")
	} else if time < 20 {
		fmt.Println("Good afternoon")
	} else {
		fmt.Println("Good evening")
	}
}