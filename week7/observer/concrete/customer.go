package concrete

import "fmt"

// Customer is the concrete Observer.
type Customer struct {
	id string
}

func NewCustomer(id string) *Customer {
	return &Customer{id: id}
}

// Implement the Observer interface method:

func (c *Customer) Update(message string) {
	fmt.Printf("Customer %s received notification: %s\n", c.id, message)
}
