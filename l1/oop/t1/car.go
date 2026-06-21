package vehicles

import "fmt"

type Car struct {
	Brand    string
	engineOn bool
}

func (c *Car) StartEngine() error {
	if c.engineOn {
		return ErrEngineAlreadyRunning
	}
	c.engineOn = true
	return nil
}

func (c *Car) StopEngine() error {
	if !c.engineOn {
		return ErrEngineOff
	}
	c.engineOn = false
	return nil
}

func (c Car) EngineStatus() bool {
	return c.engineOn
}

func (c Car) GetInfo() string {
	var engineState string
	if c.engineOn {
		engineState = "active"
	} else {
		engineState = "inactive"
	}
	return fmt.Sprintf("Brand is %q. The engine is %s.", c.Brand, engineState)
}

func (c Car) Honk() string {
	return "Beep beep!"
}

func NewCar(brand string) Car {
	return Car{
		engineOn: false,
		Brand:    brand,
	}
}
