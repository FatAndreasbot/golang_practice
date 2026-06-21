package users

type Admin struct {
	Moderator
}

func NewAdmin(username string) Admin {
	base := NewModerator(username)
	base.permissions = append(base.permissions, Delete)
	base.permissions = append(base.permissions, ManageRoles)
	return Admin{
		Moderator: base,
	}
}
