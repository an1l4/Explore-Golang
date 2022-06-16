package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		//fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)

	total := 1
	for total < 1000 { //while
		total += total
	}
	fmt.Println(total)

}