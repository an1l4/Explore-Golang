package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	m["Answer"] = 42
	fmt.Println(m)

	m["Answer"] = 48
	fmt.Println(m)

	delete(m, "Answer")
	fmt.Println("The value is", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value is ", v, "Present?", ok)

}
