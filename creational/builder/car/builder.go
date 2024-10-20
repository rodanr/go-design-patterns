package car

func (c *Builder) SetWheels() IBuilder {
	fourWheels := 4
	c.car.Wheels = fourWheels
	return c
}

func (c *Builder) SetSeats(seats int) IBuilder {
	c.car.Seats = seats
	return c
}

func (c *Builder) SetEngine(engine Engine) IBuilder {
	c.car.Engine = engine
	return c
}

func (c *Builder) SetColor(color Color) IBuilder {
	c.car.Color = color
	return c
}

func (c *Builder) SetType(t Type) IBuilder {
	c.car.Type = t
	return c
}

func (c *Builder) Build() Car {
	const (
		defaultWheels = 4
		defaultSeats  = 4
		defaultEngine = EngineV8
		defaultColor  = ColorRed
		defaultType   = TypeSedan
	)

	if c.car.Wheels == 0 {
		c.car.Wheels = defaultWheels
	}

	if c.car.Seats == 0 {
		c.car.Seats = defaultSeats
	}

	if c.car.Engine == "" {
		c.car.Engine = defaultEngine
	}

	if c.car.Color == "" {
		c.car.Color = defaultColor
	}

	if c.car.Type == "" {
		c.car.Type = defaultType
	}

	return c.car
}

func NewBuilder() IBuilder {
	return &Builder{}
}
