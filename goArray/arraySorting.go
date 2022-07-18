package main

import "fmt"

func main() {
	var array = [100]int{}
	var limit, temp int

	fmt.Println("enter the limit")
	fmt.Scanf("%d", &limit)

	fmt.Println("enter the values")
	for i := 0; i < limit; i++ {
		fmt.Scanf("%d", &array[i])

	}
	for i := 0; i < limit-1; i++ {
		for j := i + 1; j < limit; j++ {
			if array[i] > array[j] {
				temp = array[i]
				array[i] = array[j]
				array[j] = temp

			}
		}
	}
	fmt.Println("sorted array")
	for i := 0; i < limit; i++ {
		fmt.Println(array[i])

	}

}
