package main

import (
	"fmt"
)

type triangle struct{
	height float64
	base float64
}

type square struct{
	sideLength float64
}

type shape interface{
	getArea() float64
}

func main() {
	sq := square{10}
	tri := triangle{10, 10}

	printArea(sq)
	printArea(tri)
}

func (t triangle) getArea() float64 {
	return (t.height * t.base)/2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
 fmt.Println(s.getArea())
}

