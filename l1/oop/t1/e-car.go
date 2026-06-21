package vehicles

type ElectricCar struct {
	Car
	batteryLevel int
}

func (e *ElectricCar) StartEngine() error {
	if e.engineOn {
		return ErrEngineAlreadyRunning
	}
	if e.batteryLevel < 5 {
		return ErrLowBattery
	}
	e.engineOn = true
	e.batteryLevel -= 10

	return nil
}

func (e ElectricCar) GetBatteryLevel() int {
	return e.batteryLevel
}

func NewECar(brand string) ElectricCar {
	return ElectricCar{
		Car:          NewCar(brand),
		batteryLevel: 100,
	}
}
