package main

import "fmt"

func plus(x int,y int) int {
	return x + y

}

func plusPlus(x,y,z int)int {
	return x + y + z
}

func main() {
	result := plus(3,3)
	result1 := plusPlus(5,5,5)

	fmt.Println(result)
	fmt.Println(result1)
}