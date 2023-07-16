package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	hight float64
	width float64
}

func (r Rectangle) Area() float64 {
	return r.hight * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.hight)
}
func main() {
	var s Shape
	s = Rectangle{width: 7.5, hight: 5.6}
	area := s.Area()
	peri := s.Perimeter()

	fmt.Println("The area and the perimete are as follows", area, peri)
}
