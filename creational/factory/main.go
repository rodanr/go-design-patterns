package main

import "fmt"

type AnimalType int

const (
	AnimalTypeDog  AnimalType = iota
	AnimalTypeCat  AnimalType = iota
	AnimalTypeFish AnimalType = iota
)

type Animal interface {
	Speak() string
	Breathe() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Breathe() string {
	return "Breathing through lungs"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Breathe() string {
	return "Breathing through lungs"
}

type Fish struct{}

func (f Fish) Speak() string {
	return "..."
}

func (f Fish) Breathe() string {
	return "Breathing through gills"
}

func NewAnimal(animalType AnimalType) Animal {
	switch animalType {
	case AnimalTypeDog:
		return Dog{}
	case AnimalTypeCat:
		return Cat{}
	case AnimalTypeFish:
		return Fish{}
	default:
		return Dog{}
	}
}

func main() {
	dog := NewAnimal(AnimalTypeDog)
	cat := NewAnimal(AnimalTypeCat)
	fish := NewAnimal(AnimalTypeFish)

	animals := []Animal{dog, cat, fish}

	for _, animal := range animals {
		fmt.Printf("Animal says: %s\n", animal.Speak())
	}
}
