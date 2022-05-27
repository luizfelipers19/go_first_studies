package main

import "fmt"

type shape interface {
	getArea() float64
}

func main() {

	tr := triangle{
		sides:  3,
		height: 10,
		base:   5}

	sq := square{
		sides: 4,
		base:  4,
	}

	printArea(tr)
	printArea(sq)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

type triangle struct {
	sides  int
	height float64
	base   float64
}

type square struct {
	sides int
	base  float64
}

func (sq square) getArea() float64 {
	return sq.base * sq.base
}

func (tr triangle) getArea() float64 {
	return (tr.height * tr.base) / 2
}
