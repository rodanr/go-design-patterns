package main

import "fmt"

type Car interface {
	Drive() string
}

type Engine interface {
	Start() string
}

type SportsCar struct{}

func (s *SportsCar) Drive() string {
	return "Driving a sports car"
}

type SportsEngine struct{}

func (e *SportsEngine) Start() string {
	return "Starting a sports engine"
}

type LuxuryCar struct{}

func (l *LuxuryCar) Drive() string {
	return "Driving a luxury car"
}

type LuxuryEngine struct{}

func (l *LuxuryEngine) Start() string {
	return "Starting a luxury engine"
}

type CarFactory interface {
	MakeCar() Car
	MakeEngine() Engine
}

type SportsCarFactory struct{}

func (s *SportsCarFactory) MakeCar() Car {
	return &SportsCar{}
}

func (s *SportsCarFactory) MakeEngine() Engine {
	return &SportsEngine{}
}

type LuxuryCarFactory struct{}

func (l *LuxuryCarFactory) MakeCar() Car {
	return &LuxuryCar{}
}

func (l *LuxuryCarFactory) MakeEngine() Engine {
	return &LuxuryEngine{}
}

func main() {
	var carFactory CarFactory

	// You can dynamically switch between factories
	carFactory = &SportsCarFactory{}
	// carFactory = &LuxuryCarFactory{} // Uncomment this to switch to luxury car

	car := carFactory.MakeCar()
	engine := carFactory.MakeEngine()

	fmt.Println(car.Drive())
	fmt.Println(engine.Start())
}
