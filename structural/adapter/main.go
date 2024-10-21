package main

import "fmt"

type Car interface {
	Drive() string
}

type ElectricCar struct{}

func (e *ElectricCar) ChargeAndGo() string {
	return "Driving an electric car after charging"
}

type GasolineCar struct{}

func (g *GasolineCar) FillAndGo() string {
	return "Driving a gasoline car after filling the tank"
}

type ElectricCarAdapter struct {
	ElectricCar *ElectricCar
}

func (e *ElectricCarAdapter) Drive() string {
	return e.ElectricCar.ChargeAndGo()
}

type GasolineCarAdapter struct {
	GasolineCar *GasolineCar
}

func (g *GasolineCarAdapter) Drive() string {
	return g.GasolineCar.FillAndGo()
}

func main() {
	// Create instances of the adaptees
	electricCar := &ElectricCar{}
	gasolineCar := &GasolineCar{}

	// Wrap the adaptees in their respective adapters
	electricCarAdapter := &ElectricCarAdapter{ElectricCar: electricCar}
	gasolineCarAdapter := &GasolineCarAdapter{GasolineCar: gasolineCar}

	fmt.Println(electricCarAdapter.Drive())
	fmt.Println(gasolineCarAdapter.Drive())
}
