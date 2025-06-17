package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	var s Shape
	s = Circle{5}
	fmt.Printf("Area of Circle: %.2f \n", s.Area())
	s = Rectangle{5, 10}
	fmt.Println("Area of Rectangle:", s.Area())
}
