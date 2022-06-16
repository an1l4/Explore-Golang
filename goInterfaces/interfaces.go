package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perm() float64
}

type react struct {
	width,height float64
}

type circle struct {
	radius float64
}

func (r react) area() float64{
	return r.width * r.height

}
func (r react) perm() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perm() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
fmt.Println(g)
fmt.Println(g.area())
fmt.Println(g.perm())
}

func main() {
	r := react{width:3,height:4}
	c := circle{radius:5}

	measure(r)
	measure(c)
}