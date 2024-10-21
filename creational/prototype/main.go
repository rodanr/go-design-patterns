package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Color string
}

func (c *Car) Clone() *Car {
	return &Car{
		Brand: c.Brand,
		Model: c.Model,
		Color: c.Color,
	}
}

func main() {
	originalCar := &Car{
		Brand: "Tesla",
		Model: "Model 3",
		Color: "Red",
	}

	clonedCar := originalCar.Clone()

	clonedCar.Color = "Blue"

	fmt.Printf("Original Car: %+v\n", originalCar)
	fmt.Printf("Cloned Car: %+v\n", clonedCar)
}
