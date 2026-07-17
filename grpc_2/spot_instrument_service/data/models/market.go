package models

import (
	"proto/user_service"
	"time"

	"github.com/google/uuid"
)

type Market struct {
	Name         string
	ID           uuid.UUID
	Enabled      bool
	DeletedAt    *time.Time
	AllowedRoles map[user_service.UserRole]struct{}
}

func NewMarket(name string, allowedRoles ...user_service.UserRole) *Market {
	result := Market{
		ID:           uuid.New(),
		Name:         name,
		AllowedRoles: make(map[user_service.UserRole]struct{}),
		Enabled:      true,
		DeletedAt:    nil,
	}
	for _, role := range allowedRoles {
		result.AllowedRoles[role] = struct{}{}
	}

	return &result
}

func (m *Market) Toggle() {
	m.Enabled = !m.Enabled
}

func (m *Market) Delete() {
	if m.DeletedAt == nil {
		now := time.Now()
		m.DeletedAt = &now
	}
}

func (m *Market) IsValid() bool {
	return m.Enabled && m.DeletedAt == nil
}
