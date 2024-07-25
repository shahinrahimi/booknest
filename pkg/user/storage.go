package user

type Storage interface {
	GetUsers() ([]*User, error)
	GetUser(id string) (*User, error)
	CreateUser(u User) error
	UpdateUser(id string, u User) error
	DeleteUser(id string) error
}
