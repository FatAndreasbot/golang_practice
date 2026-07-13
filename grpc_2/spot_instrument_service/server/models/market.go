package models

import (
	"errors"
	"fmt"
	"time"
)

type Role int

const (
	ADMIN Role = iota
	BROKER
	USER
)

func GetRoleFromName(name string) (Role, error) {
	switch name {
	case "ADMIN":
		return ADMIN, nil
	case "BROKER":
		return BROKER, nil
	case "USER":
		return USER, nil
	default:
		return -1, fmt.Errorf("role %q does not exists", name)
	}
}

func (r Role) ToString() (string, error) {
	switch r {
	case ADMIN:
		return "ADMIN", nil
	case BROKER:
		return "BROKER", nil
	case USER:
		return "USER", nil
	default:
		return "", errors.New("role value was not found")
	}
}

type Market struct {
	Name         string
	Enabled      bool
	DeletedAt    *time.Time
	AllowedRoles map[Role]struct{}
}

func NewMarket(name string, allowedRoles []Role) *Market {
	result := Market{
		Name:         name,
		AllowedRoles: make(map[Role]struct{}),
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
