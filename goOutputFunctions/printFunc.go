package main

import "fmt"

func main() {
	var i,j string = "Hello","World"
	var x,y int = 10,20
	fmt.Print(i)
	fmt.Print(j,"\n")

	//new line
	fmt.Print(i,"\n")
	fmt.Print(j,"\n")

	//space in between
	fmt.Print(i," ",j,"\n")

	//numbers-automatically insert space if its not string
	fmt.Print(x,y)



}