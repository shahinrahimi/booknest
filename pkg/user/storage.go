package user

type Storage interface {
	GetUsers() []*User
	GetUser(id string) *User
	CreateUser() error
	UpdateUser() error
	DeleteUser() error
}
