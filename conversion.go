package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our channel")
	fmt.Println("please rate between 1 to 5")

	reader := bufio.NewReader(os.Stdin) //want to read something from user input, reader is a reference here

	input, _ := reader.ReadString('\n') //reader reads the input until i hit enter key

	fmt.Println("Thanks for rating us", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //converting string to float type and removing spaces

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to your rating", numRating+1)
	}
}
