package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
//continue stmt is used to skip one or more iterations in the loop