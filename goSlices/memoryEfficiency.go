package main

import "fmt"
	//copy function to create required elements
	//syntx copy(dest,src)
	func main() {
		numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
		fmt.Println(numbers)
	
	//create copy
	neededNumbers := numbers[:len(numbers)-10]
	fmt.Println(neededNumbers)

	numberSliceCopy := make([]int,len(neededNumbers))
	fmt.Println(numberSliceCopy)
	copy(numberSliceCopy,neededNumbers)
	fmt.Println(numberSliceCopy)
	

	}

	