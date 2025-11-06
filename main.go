package main

import (
	. "SoftwareDesignPatterns/week8"
	"fmt"
)

func main() {
	// Declaring elements
	circle := &Circle{Radius: 5}
	square := &Square{Side: 4}

	shapes := []Shape{circle, square}

	// 1. Using AreaCalculator
	fmt.Println("--- Using AreaCalculator ---")
	areaVisitor := &AreaCalculator{}

	for _, shape := range shapes {
		shape.Accept(areaVisitor)
	}

	// 2. Using Draw
	fmt.Println("\n--- Using Draw Visitor ---")
	drawVisitor := &Draw{}

	for _, shape := range shapes {
		shape.Accept(drawVisitor)
	}
}
