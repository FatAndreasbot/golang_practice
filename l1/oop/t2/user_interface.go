package users

type User interface {
	GetUsername() string
	HasPermission(permission string) bool
	GetRole() string
}
