package concrete

import "fmt"

// PercentageDiscount (percentage discount)
type PercentageDiscount struct {
	Percentage float64 // Percent of discount
}

func (s *PercentageDiscount) ApplyDiscount(total float64) float64 {
	discount := total * s.Percentage
	finalAmount := total - discount
	fmt.Printf("-> Applied: Discount %.0f%%. Saved: %.2f.\n", s.Percentage*100, discount)
	return finalAmount
}
