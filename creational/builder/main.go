package main

import (
	"builder/car"
	"fmt"
)

func main() {
	carBuilder := car.NewBuilder()
	car := carBuilder.SetWheels().SetSeats(4).SetEngine(car.EngineV8).SetColor(car.ColorRed).SetType(car.TypeSedan).Build()
	fmt.Printf("Car: %+v\n", car)
}
