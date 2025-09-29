package main

import (
	. "SoftwareDesignPatterns/week3"
	"fmt"
)

func main() {
	// creating adaptee
	stringService := &StringService{}

	// creating adapter with adaptee
	calculatorAdapter := &ServiceAdapter{
		Service: stringService,
	}

	// using adapter and expecting int output
	result, err := calculatorAdapter.Add(100, 25)

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Успешно! 100 + 25 = %d (Тип: %T)\n", result, result)
	}
}
