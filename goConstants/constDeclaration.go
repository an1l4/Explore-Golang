package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	const PI = 3.14
	fmt.Println(PI)

	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	//using context
	fmt.Println(math.Sin(n))
}