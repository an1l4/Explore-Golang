package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitList)
	fmt.Printf("type of fruitList %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	highscores := make([]int, 4)

	highscores[0] = 245
	highscores[1] = 743
	highscores[2] = 549
	highscores[3] = 852
	fmt.Println(highscores)

	highscores = append(highscores, 321, 964, 459)
	fmt.Println(highscores)

	//sorting the values
	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println(highscores)

	//checking values are sorted or not
	fmt.Println(sort.IntsAreSorted(highscores))
}
