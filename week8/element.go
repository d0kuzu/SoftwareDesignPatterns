package week8

// Shape (Element)
type Shape interface {
	Accept(Visitor)
	getType() string
}

// Circle concrete element
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c) // Calls specific method for circle
}

func (c *Circle) getType() string {
	return "Circle"
}

// Square concrete element
type Square struct {
	Side float64
}

func (s *Square) Accept(v Visitor) {
	v.VisitSquare(s) // Calls specific method for square
}

func (s *Square) getType() string {
	return "Square"
}
