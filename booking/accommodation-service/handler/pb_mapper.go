package handler

import (
	"accommodation-service/model"
	accommodation "common/proto/accommodation-service/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func mapAccommodation(modelAccommodation *model.Accommodation) *accommodation.Accommodation {

	var pbAvailabilities []*accommodation.Availability
	for _, availability := range modelAccommodation.Availabilities {
		pbAvailabilities = append(pbAvailabilities, mapAvailabilityPb(availability))
	}

	accommodationPb := &accommodation.Accommodation{
		Id:                modelAccommodation.Id.Hex(),
		Availabilities:    pbAvailabilities,
		Name:              modelAccommodation.Name,
		MinNumberOfGuests: int32(modelAccommodation.MinNumberOfGuests),
		MaxNumberOfGuests: int32(modelAccommodation.MaxNumberOfGuests),
		Address: &accommodation.AddressDTO{
			Country: modelAccommodation.Address.Country,
			City:    modelAccommodation.Address.City,
			Street:  modelAccommodation.Address.Street,
			Number:  modelAccommodation.Address.StreetNumber,
		},
		HostID:      modelAccommodation.HostID,
		Benefits:    modelAccommodation.Benefits,
		IsSuperHost: modelAccommodation.IsSuperHost,
	}
	return accommodationPb
}

func mapNewAccommodation(accommodationPb *accommodation.NewAccommodation) *model.Accommodation {

	var modelAvailabilities []*model.Availability
	for _, availability := range accommodationPb.Availabilities {
		modelAvailabilities = append(modelAvailabilities, mapAvailabilityPbToModel(availability))
	}
	accommodation := &model.Accommodation{
		Name:              accommodationPb.Name,
		Availabilities:    modelAvailabilities,
		MinNumberOfGuests: int(accommodationPb.MinNumberOfGuests),
		MaxNumberOfGuests: int(accommodationPb.MaxNumberOfGuests),
		Address: model.Address{
			Country:      accommodationPb.Address.Country,
			City:         accommodationPb.Address.City,
			Street:       accommodationPb.Address.Street,
			StreetNumber: accommodationPb.Address.Number,
		},
		HostID:   accommodationPb.HostID,
		Benefits: accommodationPb.Benefits,
	}

	return accommodation
}

func mapUpdateAccommodationPb(accommodationPb *accommodation.Accommodation) *model.Accommodation {
	AccommodationId, _ := primitive.ObjectIDFromHex(accommodationPb.Id)

	var modelAvailabilities []*model.Availability
	for _, availability := range accommodationPb.Availabilities {
		modelAvailabilities = append(modelAvailabilities, mapAvailabilityPbToModel(availability))
	}

	accommodation := &model.Accommodation{
		Id:                AccommodationId,
		Availabilities:    modelAvailabilities,
		Name:              accommodationPb.Name,
		MinNumberOfGuests: int(accommodationPb.MinNumberOfGuests),
		MaxNumberOfGuests: int(accommodationPb.MaxNumberOfGuests),
		Address: model.Address{
			Country:      accommodationPb.Address.Country,
			City:         accommodationPb.Address.City,
			Street:       accommodationPb.Address.Street,
			StreetNumber: accommodationPb.Address.Number,
		},
		HostID:      accommodationPb.HostID,
		Benefits:    accommodationPb.Benefits,
		IsSuperHost: accommodationPb.IsSuperHost,
	}

	return accommodation
}

func mapAvailabilityPb(modelAvailability *model.Availability) *accommodation.Availability {
	return &accommodation.Availability{
		Id:             modelAvailability.Id.Hex(),
		StartDate:      modelAvailability.StartDate.Format("2006-01-02"),
		EndDate:        modelAvailability.EndDate.Format("2006-01-02"),
		Price:          float32(modelAvailability.Price),
		PriceSelection: accommodation.PriceSelection(modelAvailability.PriceSelection),
	}
}
func mapAvailabilityPbToModel(availabilityPb *accommodation.Availability) *model.Availability {
	startDate, _ := time.Parse("2006-01-02", availabilityPb.StartDate)
	endDate, _ := time.Parse("2006-01-02", availabilityPb.EndDate)
	availabilityId, _ := primitive.ObjectIDFromHex(availabilityPb.Id)

	return &model.Availability{
		Id:             availabilityId,
		StartDate:      startDate,
		EndDate:        endDate,
		Price:          float64(availabilityPb.Price),
		PriceSelection: model.PriceSelection(availabilityPb.PriceSelection),
	}
}
func mapNewAvailability(availabilityPb *accommodation.NewAvailability) *model.Availability {
	startDate, _ := time.Parse("2006-01-02", availabilityPb.StartDate)
	endDate, _ := time.Parse("2006-01-02", availabilityPb.EndDate)

	availability := &model.Availability{
		Id:             primitive.NewObjectID(),
		StartDate:      startDate,
		EndDate:        endDate,
		Price:          float64(availabilityPb.Price),
		PriceSelection: model.PriceSelection(availabilityPb.PriceSelection),
	}
	return availability
}
