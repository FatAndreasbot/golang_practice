package main

import (
	"devices"
	"fmt"
)

func main() {
	var device_list []devices.Device = make([]devices.Device, 0, 3)
	laptop, err := devices.NewLaptop(
		"windows 11",
		"MSI B14 modern",
		devices.Business,
	)
	if err == nil {
		device_list = append(device_list, &laptop)
	}

	phone, err := devices.NewSmartphone(
		"9.0",
		"Pixel 6",
		"Google",
	)
	if err == nil {
		device_list = append(device_list, &phone)
	}

	watch, err := devices.NewSmartwatch(
		"12345",
		"Samsung watch",
		"Step counter", "Time", "Pulse",
	)
	if err == nil {
		device_list = append(device_list, &watch)
	}

	for _, d := range device_list {
		fmt.Println(d.GetInfo())
	}
}
