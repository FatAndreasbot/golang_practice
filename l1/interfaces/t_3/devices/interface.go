package devices

import "errors"

type Device interface {
	UpdateOS(versionOS string) error
	GetInfo() string
}

var (
	ErrUnsupported = errors.New("обновление недоступно")
)
