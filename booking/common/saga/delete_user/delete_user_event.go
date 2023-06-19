package delete_user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Role        string
	Address     Address
	IsSuperHost bool
}
type Address struct {
	Street       string
	StreetNumber string
	City         string
	Country      string
}

type DeleteUserCommandType int8

const (
	CheckActiveReservations DeleteUserCommandType = iota
	DeleteUser
	CancelDeleteUser
	DeleteAccommodations
	UnknownCommand
)

type DeleteUserCommand struct {
	User User
	Type DeleteUserCommandType
}
type DeleteUserReplyType int8

const (
	DeleteUserAllowed DeleteUserReplyType = iota
	DeleteUserNotAllowed
	DeletedGuest
	DeletedHost
	AccommodationsDeleted
	UnknownReply
)

type DeleteUserReply struct {
	User User
	Type DeleteUserReplyType
}
