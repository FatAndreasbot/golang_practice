package users

import "errors"

type Permission int

const (
	Read Permission = iota
	Edit
	BanUser
	Delete
	ManageRoles
)

var ErrPermissionDoesNotExist error = errors.New("permission does not exist")

func GetPermission(name string) (Permission, error) {
	switch name {
	case "read":
		return Read, nil
	case "edit":
		return Edit, nil
	case "ban_user":
		return BanUser, nil
	case "delete":
		return Delete, nil
	case "manage_roles":
		return ManageRoles, nil
	default:
		return Read, ErrPermissionDoesNotExist
	}
}

func GetPermissionName(p Permission) (string, error) {
	switch p {
	case Read:
		return "read", nil
	case Edit:
		return "edit", nil
	case BanUser:
		return "ban_user", nil
	case Delete:
		return "delete", nil
	case ManageRoles:
		return "manage_roles", nil
	default:
		return "", ErrPermissionDoesNotExist
	}
}
