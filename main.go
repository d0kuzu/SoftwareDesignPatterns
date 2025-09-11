package main

import (
	"SoftwareDesignPatterns/week1"
	"fmt"
)

func main() {
	car := week1.NewCarBuilder().
		SetBrand("BMW").
		SetYear(2020).
		Build()

	fmt.Println(car.GetBrand(), car.GetYear())
}
