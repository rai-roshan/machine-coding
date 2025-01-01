package user

type User struct {
	id   uint32
	name string
}

func (u *User) GetUserId() uint32 {
	return u.id
}

func (u *User) GetUserName() string {
	return u.name
}

func NewUser(id uint32, name string) *User {
	return &User{
		id: id,
		name: name,
	}
}