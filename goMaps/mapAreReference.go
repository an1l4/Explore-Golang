package main

import "fmt"

func main() {
	var a = map[string]string{"brand":"ford","model":"mustag","year":"1964"}

	b := a

	b["year"] = "1970"
	fmt.Println(a)
	fmt.Println(b)
}