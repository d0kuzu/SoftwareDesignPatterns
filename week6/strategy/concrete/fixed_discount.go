package concrete

import "fmt"

// FixedAmountDiscount (fixed discount)
type FixedAmountDiscount struct {
	Amount float64 // Amount of discount
}

func (s *FixedAmountDiscount) ApplyDiscount(total float64) float64 {
	finalAmount := total - s.Amount
	if finalAmount < 0 {
		finalAmount = 0
	}
	fmt.Printf("-> Applied: Fixed discoun %.2f.\n", s.Amount)
	return finalAmount
}
