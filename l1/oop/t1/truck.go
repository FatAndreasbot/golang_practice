package vehicles

import "fmt"

type Truck struct {
	Car
	cargoCapacity int
}

func (t Truck) Honk() string {
	return "Honk Honk!"
}

func (t Truck) GetInfo() string {
	baseInfo := t.Car.GetInfo()
	return fmt.Sprintf("%s The capacity is %d", baseInfo, t.cargoCapacity)
}

func (t Truck) GetCargoCapacity() int {
	return t.cargoCapacity
}

func NewTruck(brand string, capacity int) Truck {
	return Truck{
		Car:           NewCar(brand),
		cargoCapacity: capacity,
	}
}
