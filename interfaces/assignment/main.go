package main

import "fmt"

type Shape interface {
	getArea() float64
}

type traingale struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	s := square{sideLength: 10}
	t := traingale{height: 20, base: 10}

	printArea(s)
	printArea(t)
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func (t traingale) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
