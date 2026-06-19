package devices

import "fmt"

type baseDevice struct {
	OSVersion string
	Model     string
}

func (b baseDevice) GetInfo() string {
	return fmt.Sprintf("Модель: %s, ОС: %s", b.Model, b.OSVersion)
}
