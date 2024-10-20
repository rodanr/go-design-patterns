package car

type Engine string

const (
	EngineV8  Engine = "V8"
	EngineV12 Engine = "V12"
)

type Color string

const (
	ColorRed   Color = "Red"
	ColorBlue  Color = "Blue"
	ColorBlack Color = "Black"
)

type Type string

const (
	TypeSedan Type = "Sedan"
	TypeSUV   Type = "SUV"
	TypeTruck Type = "Truck"
	TypeHatch Type = "Hatch"
)

type Car struct {
	Wheels int
	Seats  int
	Engine Engine
	Color  Color
	Type   Type
}

type IBuilder interface {
	SetWheels() IBuilder
	SetSeats(seats int) IBuilder
	SetEngine(engine Engine) IBuilder
	SetColor(color Color) IBuilder
	SetType(t Type) IBuilder
	Build() Car
}

type Builder struct {
	car Car
}
