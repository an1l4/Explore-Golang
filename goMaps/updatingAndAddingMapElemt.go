package main

import "fmt"

func main() {
	var a = make(map[string]string)
	a["brand"] = "ford"
	a["model"] = "mustag"
	a["year"] = "1964"
	fmt.Println(a)

	//updatingnd adding
	a["year"] = "1970"
	a["color"] = "red"
	fmt.Println(a)
}