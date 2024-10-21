package main

import "fmt"

type State interface {
	Print()
}

type Printer struct {
	currentState State
}

func (p *Printer) SetState(state State) {
	p.currentState = state
}

func (p *Printer) Print() {
	p.currentState.Print()
}

type ReadyState struct {
	printer *Printer
}

func (r *ReadyState) Print() {
	fmt.Println("Printing...")
	// After printing, the printer will be out of paper
	r.printer.SetState(&OutOfPaperState{printer: r.printer})
}

type OutOfPaperState struct {
	printer *Printer
}

func (o *OutOfPaperState) Print() {
	fmt.Println("Cannot print, out of paper!")
}

// Main function to demonstrate the State Pattern
func main() {
	// Create a new printer and set its initial state to Ready
	printer := &Printer{}
	readyState := &ReadyState{printer: printer}
	printer.SetState(readyState)

	// Try to print when the printer is in Ready state
	printer.Print() // Output: Printing...

	// Try to print when the printer is out of paper
	printer.Print() // Output: Cannot print, out of paper!
}
