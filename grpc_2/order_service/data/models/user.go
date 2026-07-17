package models

import (
	"proto/user_service"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID             `json:"uuid"`
	Username     string                `json:"username"`
	Role         user_service.UserRole `json:"role"`
}
