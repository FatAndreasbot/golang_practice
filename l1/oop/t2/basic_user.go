package users

import "slices"

type BasicUser struct {
	username    string
	permissions []Permission
}

func (u BasicUser) GetUsername() string {
	return u.username
}

func (u BasicUser) HasPermission(permission string) bool {
	permissionCode, err := GetPermission(permission)
	if err != nil {
		return false
	}
	return slices.Contains(u.permissions, permissionCode)
}

func NewBasicUser(username string) BasicUser {
	return BasicUser{
		username:    username,
		permissions: []Permission{Read},
	}
}

func (u *BasicUser) GetRole() string {
	return "BasicUser"
}
