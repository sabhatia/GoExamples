package main

type square struct {
	side float64
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}

type shape interface {
	getArea() float64
}

func printArea(s shape) float64 {
	return s.getArea()
}
