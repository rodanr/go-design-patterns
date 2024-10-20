package main

type Coffee interface {
	Cost() int
}

type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() int {
	return 10
}

type MilkCoffee struct {
	Coffee Coffee
}

func (m *MilkCoffee) Cost() int {
	return m.Coffee.Cost() + 10
}

type WhipCoffee struct {
	Coffee Coffee
}

func (w *WhipCoffee) Cost() int {
	return w.Coffee.Cost() + 5
}

func main() {
	simpleCoffee := &SimpleCoffee{}
	milkCoffee := &MilkCoffee{Coffee: simpleCoffee}
	whipCoffee := &WhipCoffee{Coffee: milkCoffee}

	println(whipCoffee.Cost())
}
