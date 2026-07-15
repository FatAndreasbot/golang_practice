package models

import (
	"crypto/sha256"
	"proto/user_service"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID             `json:"uuid"`
	Username     string                `json:"username"`
	passwordHash [32]byte              `json:"-"`
	Role         user_service.UserRole `json:"role"`
}

func NewUser(username, password string, role user_service.UserRole) *User{
	return &User{
		ID: uuid.New(),
		Username: username,
		passwordHash: sha256.Sum256([]byte(password)),
		Role: role,
	}
}

func (u *User) CheckPassword(pass string) bool {
	hash := sha256.Sum256([]byte(pass))
	return hash == u.passwordHash
}
