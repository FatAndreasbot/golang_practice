package devices

type Smartwatch struct {
	baseDevice
	Features []string
}

func (s *Smartwatch) UpdateOS(versionOS string) error {
	if len(versionOS) != 5 {
		return ErrUnsupported
	}

	s.OSVersion = versionOS
	return nil
}

func NewSmartwatch(osVersion string, model string, features ...string) (Smartwatch, error) {
	if len(osVersion) != 5 {
		return Smartwatch{}, ErrUnsupported
	}

	return Smartwatch{
		baseDevice: baseDevice{
			OSVersion: osVersion,
			Model:     model,
		},
		Features: features,
	}, nil
}
