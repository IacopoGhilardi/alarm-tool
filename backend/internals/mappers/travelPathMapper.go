package mappers

import (
	"github.com/iacopoghilardi/alarm-tool/internals/models"
	"github.com/iacopoghilardi/alarm-tool/internals/types/dto"
)

func CreateTravelPathDtoToTravelPathModel(travelPath *dto.CreateTravelPathDto) *models.TravelPath {
	return &models.TravelPath{
		UserID: travelPath.UserID,
		StartLocation: models.Location{
			Lat: travelPath.StartLocation.Latitude,
			Lng: travelPath.StartLocation.Longitude,
		},
		EndLocation: models.Location{
			Lat: travelPath.EndLocation.Latitude,
			Lng: travelPath.EndLocation.Longitude,
		},
	}
}
