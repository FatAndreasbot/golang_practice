package models

import (
	"crypto/sha256"
	"proto/user_service"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID
	Username string
	passwordHash [32]byte
	Role user_service.UserRole
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
