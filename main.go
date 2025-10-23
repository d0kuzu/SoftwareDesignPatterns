package main

import (
	. "SoftwareDesignPatterns/week7"
	. "SoftwareDesignPatterns/week7/strategy/concrete"
	"fmt"
)

func main() {
	cartTotal := 150.0 // Simulating total prce of cart

	// 1. Creating Context (cart), with initial strategy "NoDiscount"
	cart := &ShoppingCart{
		TotalAmount: cartTotal,
		Strategy:    &NoDiscount{}, // Инициализируем стратегией по умолчанию
	}

	fmt.Println("--- Step 1: Using default strategy (No Discount Strategy) ---")
	cart.Checkout()

	// 2. Dynamically changing out strategy to "Percentage Discount" with 15%
	fmt.Println("--- Step 2: Changing strategy to Percentage (15%) ---")
	percentStrategy := &PercentageDiscount{Percentage: 0.15}
	cart.SetStrategy(percentStrategy)
	cart.Checkout()

	// 3. Dynamically changing out strategy to "Fixed Discount" with 20$
	fmt.Println("--- Step 2: Changing strategy to Fixed (20$) ---")
	fixedStrategy := &FixedAmountDiscount{Amount: 20.0}
	cart.SetStrategy(fixedStrategy)
	cart.Checkout()
}
