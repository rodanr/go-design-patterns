package main

import "fmt"

type Command interface {
	Execute()
}

type Light struct {
	isOn bool
}

func (l *Light) On() {
	l.isOn = true
	fmt.Println("The light is now ON.")
}

func (l *Light) Off() {
	l.isOn = false
	fmt.Println("The light is now OFF.")
}

type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := &Light{}

	lightOnCommand := &LightOnCommand{light: light}
	lightOffCommand := &LightOffCommand{light: light}

	remote := &RemoteControl{}

	remote.command = lightOnCommand
	remote.PressButton()

	remote.command = lightOffCommand
	remote.PressButton()
}
