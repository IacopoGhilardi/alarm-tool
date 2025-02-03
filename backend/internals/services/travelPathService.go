package services

import (
	"github.com/iacopoghilardi/alarm-tool/internals/mappers"
	"github.com/iacopoghilardi/alarm-tool/internals/models"
	"github.com/iacopoghilardi/alarm-tool/internals/repositories"
	"github.com/iacopoghilardi/alarm-tool/internals/types/dto"
)

type TravelPathService struct {
	travelPathRepository *repositories.TravelPathRepository
}

func NewTravelPathService(travelPathRepository *repositories.TravelPathRepository) *TravelPathService {
	return &TravelPathService{travelPathRepository}
}

func (r *TravelPathService) FindAll(userId string) ([]models.TravelPath, error) {
	return r.travelPathRepository.FindAll(userId)
}

func (r *TravelPathService) Create(travelPath *dto.CreateTravelPathDto) (*models.TravelPath, error) {
	travelPathModel := mappers.CreateTravelPathDtoToTravelPathModel(travelPath)
	return r.travelPathRepository.Create(travelPathModel)
}

func (r *TravelPathService) Update(id string, travelPath *dto.CreateTravelPathDto) (*models.TravelPath, error) {
	travelPathModel := mappers.CreateTravelPathDtoToTravelPathModel(travelPath)
	return r.travelPathRepository.Update(id, travelPathModel)
}

func (r *TravelPathService) Delete(id string) error {
	return r.travelPathRepository.Delete(id)
}
