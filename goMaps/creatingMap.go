package main

import "fmt"

func main() {
	var a = map[string]string{"Brand":"Ford","Model" : "mustag"}
	b := map[string]int{"oslo" : 1,"Bergen" : 2}
	fmt.Printf("a \t %v\n",a)
	fmt.Printf("b \t %v\n",b)
}