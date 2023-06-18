package handler

import (
	events "common/saga/delete_user"
	saga "common/saga/messaging"
	"reservation-service/service"
)

type DeleteUserCommandHandler struct {
	reservationService *service.ReservationService
	replyPublisher     saga.Publisher
	commandSubscriber  saga.Subscriber
}

func NewDeleteUserCommandHandler(reservationService *service.ReservationService, publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	o := &DeleteUserCommandHandler{
		reservationService: reservationService,
		replyPublisher:     publisher,
		commandSubscriber:  subscriber,
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
	case events.CheckActiveReservations:
		activeReservations, err := handler.reservationService.GetActiveReservationsByUserId(command.User)
		if err != nil {
			return
		}
		if activeReservations != nil {
			reply.Type = events.DeleteUserNotAllowed
		} else {
			reply.Type = events.DeleteUserAllowed
		}

	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
