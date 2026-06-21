package vehicles_test

import (
	"testing"
	v "vehicles"
)

func TestVehicleStarts(t *testing.T) {
	car := v.NewCar("audi")
	eCar := v.NewECar("tesla")
	truck := v.NewTruck("kamaz", 120)

	vehicles := [3]v.Vehicle{
		&car,
		&eCar,
		&truck,
	}

	for i, vehicle := range vehicles {
		VehicleStartEngineTest(t, vehicle)
		t.Logf("tested vehicle #%d", i)
	}
}

func TestOverrides(t *testing.T) {
	car := v.NewCar("audi")
	truck := v.NewTruck("kamaz", 120)

	cHonk := car.Honk()
	tHonk := truck.Honk()

	if cHonk == tHonk {
		t.Errorf("expected different 'honks'")
	}
}

func VehicleStartEngineTest(t *testing.T, vehicle v.Vehicle) {

	err := vehicle.StopEngine()
	if err != v.ErrEngineOff {
		t.Errorf("expected error. got nil")
	}

	err = vehicle.StartEngine()
	if err != nil {
		t.Errorf("got error: %q", err.Error())
	}

	status := vehicle.EngineStatus()
	if !status {
		t.Errorf("car sohould be on")
	}
	err = vehicle.StartEngine()
	if err != v.ErrEngineAlreadyRunning {
		t.Errorf("got wrong error, or no error")
	}

	err = vehicle.StopEngine()
	if err != nil {
		t.Errorf("got error: %q", err.Error())
	}
}
