package main

import "fmt"

func main() {
	day := 5
	switch day {
	case 1,3,5,7 :
		fmt.Println("Odd day")
	case 2,4,6 :
		fmt.Println("even day")
		default :
		fmt.Println("Not a day")
	}
}