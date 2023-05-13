package model

type UserStore interface {
	GetAll() (Users, error)
	Insert(user *User) error
	DeleteAll()
	FindUserByEmail(email string) (*User, error)
	Delete(id string) error
}
