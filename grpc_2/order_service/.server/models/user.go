package models

import (
	spotInstrumentService "proto/spot_instrument_service"
)

type User struct {
	Name string
	Role spotInstrumentService.UserRole
}
