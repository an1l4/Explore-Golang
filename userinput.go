package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to our channel"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our channel")

	//what reader is gonna read ,i want to store that into a variable
	//using comma,ok syntax to store that data

	input, _ := reader.ReadString('\n') //want to read till press the enter
	fmt.Println("Thanks for rating us ", input)
	fmt.Printf("input type is %T \n", input)
}
