package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserStore interface {
	GetAll() (Users, error)
	Get(id primitive.ObjectID) (*User, error)
	Insert(user *User) error
	DeleteAll()
	FindUserByEmail(email string) (*User, error)
	Delete(id string) error
}
