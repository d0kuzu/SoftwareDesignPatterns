package week7

import (
	"SoftwareDesignPatterns/week7/strategy"
	"fmt"
)

// ShoppingCart — uses strategy.
type ShoppingCart struct {
	TotalAmount float64
	Strategy    strategy.DiscountStrategy // Field of Strategy variable
}

// SetStrategy позволяет динамически менять стратегию (поведение) во время выполнения.
func (c *ShoppingCart) SetStrategy(s strategy.DiscountStrategy) {
	c.Strategy = s
}

// Checkout delegates calculating of selected discount.
func (c *ShoppingCart) Checkout() float64 {
	fmt.Printf("Initial price: %.2f\n", c.TotalAmount)
	// Delegating: Context calls current strategy's method
	finalPrice := c.Strategy.ApplyDiscount(c.TotalAmount)
	fmt.Printf("Final price: %.2f\n\n", finalPrice)
	return finalPrice
}
