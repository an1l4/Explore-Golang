package main

import "fmt"

func main() {
	var array = [3][3]int{}
	fmt.Println("enter values")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &array[i][j])
		}
	}

	fmt.Println("entered values are")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", array[i][j])
		}
		fmt.Println()
	}

}
