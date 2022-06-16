package main

import "fmt"

func main() {
	var a = map[string]string{"brand":"ford","model":"mustag","year":"1970","day":""}

	value1,ok1 := a["brand"]
	value2,ok2 := a["color"]
	value3,ok3 := a["day"]
	_,ok4 := a["model"]

	fmt.Println(value1,ok1)
	fmt.Println(value2,ok2)
	fmt.Println(value3,ok3)
	fmt.Println(ok4)
}