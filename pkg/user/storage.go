package user

type Storage interface {
	GetUser(username string) (*User, error)
	GetUserByID(id string) (*User, error)
	GetUsers() ([]*User, error)
	CreateUser(u User) error
	UpdateUser(id string, u User) error
	DeleteUser(id string) error
}
