package main

import "fmt"

type Vertex struct{
	X int
	Y int
}

func main() {
	v := Vertex{1,2}
	v.X = 40
	fmt.Println(v.X) //field accessed using dot
	fmt.Println(v)
}