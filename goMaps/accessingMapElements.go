package main

import "fmt"

func main() {
	var a = make(map[string]string)
	a["brand"]  = "ford"
	a["model"] = "mustag"

	fmt.Println(a["brand"])
}