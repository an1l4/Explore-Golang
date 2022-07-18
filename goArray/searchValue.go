package main

import "fmt"

func main() {
	var limit, searchKey, i int
	var array = [100]int{}
	var flag int = 0

	fmt.Println("enter the limit")
	fmt.Scanf("%d", &limit)

	fmt.Println("enter values")
	for i = 0; i < limit; i++ {
		fmt.Scanf("%d", &array[i])
	}

	fmt.Println("enter the value you want to check:")
	fmt.Scanf("%d", &searchKey)

	for i = 0; i < 5; i++ {
		if searchKey == array[i] {
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Printf("found value at %d position \n", i+1)
	} else {
		fmt.Println("sorry, no match")
	}

}
