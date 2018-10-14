package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	sq := square{sideLength: 10}
	tr := triangle{base: 10, height: 10}
	printArea(sq)
	printArea(tr)
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
