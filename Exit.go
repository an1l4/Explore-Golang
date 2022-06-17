package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Hello Go")

	os.Exit(3)
}

//type ./filename(./Exit) then echo $? terminal will give u the value only
