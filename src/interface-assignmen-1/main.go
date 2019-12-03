package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type tringle struct {
	height float64
	base   float64
}

func main() {
	fmt.Println("Welcome to assignment no 1")

	ss := square{
		sideLength: 20,
	}

	printArea(ss)

	ts := tringle{
		height: 20,
		base:   30,
	}

	printArea(ts)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t tringle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(s shape) {
	fmt.Println(s.getArea())

}
