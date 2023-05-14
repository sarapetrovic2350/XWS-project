package handler

import (
	"accommodation-service/model"
	accommodation "common/proto/accommodation-service/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		HostID:   modelAccommodation.HostID,
		Benefits: modelAccommodation.Benefits,
	}
	return accommodationPb
}

func mapNewAccommodation(accommodationPb *accommodation.NewAccommodation) *model.Accommodation {

	accommodation := &model.Accommodation{

		Id:                primitive.NewObjectID(),
		Name:              accommodationPb.Name,
		MinNumberOfGuests: int(accommodationPb.MinNumberOfGuests),
		MaxNumberOfGuests: int(accommodationPb.MaxNumberOfGuests),
		Address: model.Address{
			Country:      accommodationPb.Address.Country,
			City:         accommodationPb.Address.City,
			Street:       accommodationPb.Address.Street,
			StreetNumber: accommodationPb.Address.Street,
		},
		HostID:   accommodationPb.HostID,
		Benefits: accommodationPb.Benefits,
	}
	return accommodation
}

func mapAvailabilityPb(modelAvailability *model.Availability) *accommodation.Availability {
	return &accommodation.Availability{
		Id:             modelAvailability.Id.Hex(),
		StartDate:      timestamppb.New(modelAvailability.StartDate),
		EndDate:        timestamppb.New(modelAvailability.EndDate),
		Price:          float32(modelAvailability.Price),
		PriceSelection: accommodation.PriceSelection(modelAvailability.PriceSelection),
	}
}
