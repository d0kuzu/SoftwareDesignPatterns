package week1

type Car struct {
	brand string
	year  int
}

func (c *Car) GetBrand() string { return c.brand }

func (c *Car) GetYear() int { return c.year }
