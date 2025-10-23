package strategy

// DiscountStrategy declares interface, that realises all concrete strategies.
type DiscountStrategy interface {
	// ApplyDiscount applies discount to total price.
	ApplyDiscount(total float64) float64
}
