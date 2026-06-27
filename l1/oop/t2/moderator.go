package users

type Moderator struct {
	BasicUser
}

func NewModerator(username string) Moderator {
	base := NewBasicUser(username)
	base.permissions = append(base.permissions, Edit)
	base.permissions = append(base.permissions, BanUser)
	return Moderator{
		BasicUser: base,
	}
}

func (u *Moderator) GetRole() string {
	return "Moderator"
}
