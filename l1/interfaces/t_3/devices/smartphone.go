package devices

import "strconv"

type Smartphone struct {
	baseDevice
	Manufacturer string
}

func (s *Smartphone) UpdateOS(versionOS string) error {
	versionNumber, err := strconv.ParseFloat(versionOS, 64)
	if err != nil {
		return ErrUnsupported
	}

	if versionNumber > 12.0 {
		return ErrUnsupported
	}

	s.OSVersion = versionOS
	return nil
}

func NewSmartphone(osVersion string, model string, manufacturer string) (Smartphone, error) {
	_, err := strconv.ParseFloat(osVersion, 64)
	if err != nil {
		return Smartphone{}, ErrUnsupported
	}

	return Smartphone{
		baseDevice: baseDevice{
			OSVersion: osVersion,
			Model:     model,
		},
		Manufacturer: manufacturer,
	}, nil
}
