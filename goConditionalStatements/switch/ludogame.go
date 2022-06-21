package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to Ludo")

	//to generate random number
	rand.Seed(time.Now().UnixNano()) // we get random number everytime

	diceNumber := rand.Intn(6) + 1 //number in range 5 add 1 to get 6 as well
	fmt.Println("the value of dice is :", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Allowed to start the game")
	case 2:
		fmt.Println("Allowed to move 2 spot")
	case 3:
		fmt.Println("Allowed to move 3 spot")
		fallthrough // it will go the next case as well here case 4
	case 4:
		fmt.Println("Allowed to move 4 spot")
	case 5:
		fmt.Println("Allowed to move 5 spot")
	case 6:
		fmt.Println("Allowed to move 6 spot and play again")
	default:
		fmt.Println("No move")
	}
}
