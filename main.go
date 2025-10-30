package main

import (
	. "SoftwareDesignPatterns/week7/observer/concrete"
	. "SoftwareDesignPatterns/week7/subject/concrete"
	"fmt"
)

func main() {
	// 1. Create a Subject (Item)
	phone := NewItem("Smartphone X")

	// 2. Create Observers (Customers)
	customer1 := NewCustomer("C1")
	customer2 := NewCustomer("C2")

	// 3. Register Observers
	phone.RegisterObserver(customer1)
	phone.RegisterObserver(customer2)

	// 4. Change Subject's state -> Notification
	phone.UpdateAvailability(true)
	// Output:
	// Customer C1 received notification: Smartphone X: Now in stock!
	// Customer C2 received notification: Smartphone X: Now in stock!

	fmt.Println("---")

	// 5. Change Subject's state again -> Notification
	phone.UpdateAvailability(false)
	// Output:
	// Customer C1 received notification: Smartphone X: Out of stock.
	// Customer C2 received notification: Smartphone X: Out of stock.
}
