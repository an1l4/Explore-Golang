package main

import "fmt"

func main() {
	var a = make(map[string]string)
	a["brand"] = "ford"
	a["model"] = "mustag"
	b := make(map[string]int)
	b["oslo"] = 1
	b["bergen"] = 2

	fmt.Printf("a \t %v \n",a)
	fmt.Printf("b \t %v \n",b)
}