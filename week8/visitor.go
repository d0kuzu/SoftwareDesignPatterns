package week8

import "fmt"

// Visitor interface
type Visitor interface {
	VisitCircle(*Circle)
	VisitSquare(*Square)
}

// AreaCalculator concrete visitor
type AreaCalculator struct {
	Area float64
}

func (a *AreaCalculator) VisitCircle(c *Circle) {
	// Logic to compute area for circle
	a.Area = 3.14 * c.Radius * c.Radius
	fmt.Printf("Calculating area for %s: %.2f\n", c.getType(), a.Area)
}

func (a *AreaCalculator) VisitSquare(s *Square) {
	// Logic to compute area for square
	a.Area = s.Side * s.Side
	fmt.Printf("Calculating area for %s: %.2f\n", s.getType(), a.Area)
}

// Draw concrete visitor
type Draw struct{}

func (d *Draw) VisitCircle(c *Circle) {
	fmt.Printf("Drawing a Circle with radius %.2f\n", c.Radius)
}

func (d *Draw) VisitSquare(s *Square) {
	fmt.Printf("Drawing a Square with side %.2f\n", s.Side)
}
