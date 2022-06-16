package main

import "fmt"

func main() {
	var a = make(map[string]string)
	fmt.Println(a)
	a["brand"] = "ford"
	a["model"] = "mustag"
	a["color"] = "red"
	a["year"] = "1970"
	fmt.Println(a)

	delete(a,"year")
	fmt.Println(a)
}