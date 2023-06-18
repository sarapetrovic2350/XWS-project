package handler

import (
	"accommodation-service/service"
	events "common/saga/delete_user"
	saga "common/saga/messaging"
)

type DeleteUserCommandHandler struct {
	accommodationService *service.AccommodationService
	replyPublisher       saga.Publisher
	commandSubscriber    saga.Subscriber
}

func NewDeleteUserCommandHandler(accommodationService *service.AccommodationService, publisher saga.Publisher, subscriber saga.Subscriber) (*DeleteUserCommandHandler, error) {
	o := &DeleteUserCommandHandler{
		accommodationService: accommodationService,
		replyPublisher:       publisher,
		commandSubscriber:    subscriber,
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
	case events.DeleteAccommodations:
		accommodations, err := handler.accommodationService.GetAllAccommodations()
		if err != nil {
			reply.Type = events.UnknownReply
			return
		}
		for _, accommodation := range accommodations {
			if accommodation.HostID == command.User.Id.Hex() {
				err := handler.accommodationService.Delete(accommodation.Id.Hex())
				if err != nil {
					reply.Type = events.UnknownReply
					break
				}
			}
		}
		reply.Type = events.AccommodationsDeleted
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
