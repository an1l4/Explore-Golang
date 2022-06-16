package main

import "fmt"

type react struct{
	width,height int
}

func (r *react) area() int {
	return r.width * r.height
}

func (r react) perm() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := react{width:10,height:5}
	fmt.Println(r.area())
	fmt.Println(r.perm())

	rp:=&r
	fmt.Println(rp.area())
	fmt.Println(rp.perm())

}