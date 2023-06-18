package handler

import (
	events "common/saga/delete_user"
	saga "common/saga/messaging"
	"errors"
	"user-service/service"
)

type DeleteUserCommandHandler struct {
	userService       *service.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewDeleteUserCommandHandler(userService *service.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	o := &DeleteUserCommandHandler{
		userService:       userService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *DeleteUserCommandHandler) handle(command *events.DeleteUserCommand) {

	reply := events.DeleteUserReply{User: command.User}

	switch command.Type {
	case events.DeleteUser:
		err := handler.userService.Delete(command.User.Id.Hex())
		if err != nil {
			return
		}
		if command.User.Role == "HOST" {
			reply.Type = events.DeletedHost
		} else {
			reply.Type = events.DeletedGuest
		}
	case events.CancelDeleteUser:
		errors.New("user can not be deleted")
		return

	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
