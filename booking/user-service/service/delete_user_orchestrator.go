package service

import (
	events "common/saga/delete_user"
	saga "common/saga/messaging"
	"user-service/model"
)

type DeleteUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewDeleteUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteUserOrchestrator, error) {
	o := &DeleteUserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}
func (o *DeleteUserOrchestrator) Start(user *model.User) error {
	event := &events.DeleteUserCommand{
		Type: events.CheckActiveReservations,
		User: events.User{
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  user.Password,
			Role:      user.Role,
			Address: events.Address(model.Address{
				Country:      user.Address.Country,
				City:         user.Address.City,
				Street:       user.Address.Street,
				StreetNumber: user.Address.StreetNumber,
			}),
			IsSuperHost: user.IsSuperHost,
		},
	}

	return o.commandPublisher.Publish(event)
}
func (o *DeleteUserOrchestrator) handle(reply events.DeleteUserReply) {
	command := events.DeleteUserCommand{User: reply.User}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *DeleteUserOrchestrator) nextCommandType(reply events.DeleteUserReplyType) events.DeleteUserCommandType {
	switch reply {
	case events.DeleteUserAllowed:
		return events.DeleteUser
	case events.DeleteUserNotAllowed:
		return events.CancelDeleteUser
	case events.DeletedHost:
		return events.DeleteAccommodations
	default:
		return events.UnknownCommand
	}
}
