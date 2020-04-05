package main

import "fmt"

func main() {
	// initialize the structs
	sq1 := square{10.0}
	tr1 := triangle{20.0, 30.0}

	fmt.Println("Square Side Length", sq1.side)
	fmt.Println("Area of Square", sq1.getArea())

	fmt.Println("Triagle Dimension: Base:", tr1.base, ", Height:", tr1.height)
	fmt.Println("Area of Triangle:", tr1.getArea())

	// Try the same with interfaces
	fmt.Println("Interface: Area of Square", printArea(sq1))
	fmt.Println("Interface: Area of Triangle:", printArea(tr1))
}
