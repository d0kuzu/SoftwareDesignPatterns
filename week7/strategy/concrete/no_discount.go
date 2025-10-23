package concrete

import "fmt"

// NoDiscount (without discount)
type NoDiscount struct{}

func (s *NoDiscount) ApplyDiscount(total float64) float64 {
	fmt.Println("-> Applied: No discount (full price).")
	return total
}
