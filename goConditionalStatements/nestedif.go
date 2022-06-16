package main

import "fmt"

func main() {
	num := 20
	if num >= 10 {
		fmt.Println("num is greater than 10")
		if num > 15 {
			fmt.Println("num is greater than 15 also")
		}
	} else {
		fmt.Println("num is less than 10")
	}
}