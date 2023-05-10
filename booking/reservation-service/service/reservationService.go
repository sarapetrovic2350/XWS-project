package service

import (
	"Rest/model"
	"Rest/repository"
)

type ReservationService struct {
	// NoSQL: injecting user repository
	ReservationRepo *repository.ReservationRepo
}

func NewReservationService(r *repository.ReservationRepo) *ReservationService {
	return &ReservationService{r}
}

func (service *ReservationService) CreateReservation(reservation *model.Reservation) error {
	err := service.ReservationRepo.Insert(reservation)
	if err != nil {
		return err
	}
	return nil
}

func (service *ReservationService) GetAllReservations() (model.Reservations, error) {
	reservations, err := service.ReservationRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (service *ReservationService) GetReservationsByUserId(userId string) (model.Reservations, error) {
	reservations, err := service.ReservationRepo.GetAll()
	var retReservations model.Reservations
	for _, itr := range reservations {
		if itr.UserID == userId {
			retReservations = append(retReservations, itr)
		}
	}

	if err != nil {
		return nil, err
	}
	return retReservations, nil
}
