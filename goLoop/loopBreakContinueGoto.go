package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// ways of using for loop

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, value := range days {
		fmt.Printf("index is %v and value is %v \n", index, value)
	}

	rougueValue := 1

	for rougueValue <= 10 {
		if rougueValue == 3 {
			goto wantToComeHere
		}

		if rougueValue == 5 {
			rougueValue++
			continue //if use breakit will end the loop

		}
		fmt.Println(rougueValue)
		rougueValue++
	}
	//goto statements

wantToComeHere:
	fmt.Println("come to this page")
}
