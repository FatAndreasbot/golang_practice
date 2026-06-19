package devices

import (
	"strings"
)

type LaptopType int

const (
	Gaming LaptopType = iota
	Business
	Transformer
)

type Laptop struct {
	baseDevice
	Type LaptopType
}

func (l *Laptop) UpdateOS(versionOS string) error {
	_, hasCorrectPrefix := strings.CutPrefix(versionOS, "windows")
	if !hasCorrectPrefix {
		return ErrUnsupported
	}
	l.OSVersion = versionOS

	return nil
}

func NewLaptop(osVersion string, model string, category LaptopType) (Laptop, error) {
	_, hasCorrectPrefix := strings.CutPrefix(osVersion, "windows")
	if !hasCorrectPrefix {
		return Laptop{}, ErrUnsupported
	}

	return Laptop{
		baseDevice: baseDevice{
			OSVersion: osVersion,
			Model:     model,
		},
		Type: category,
	}, nil
}
