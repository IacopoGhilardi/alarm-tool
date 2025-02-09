package services

import (
	"github.com/iacopoghilardi/alarm-tool/internals/mappers"
	"github.com/iacopoghilardi/alarm-tool/internals/models"
	"github.com/iacopoghilardi/alarm-tool/internals/repositories"
	"github.com/iacopoghilardi/alarm-tool/internals/types/dto"
)

type AlarmService struct {
	repository *repositories.AlarmRepository
}

func NewAlarmService(repository *repositories.AlarmRepository) *AlarmService {
	return &AlarmService{repository: repository}
}

func (s *AlarmService) FindAll(userId uint) ([]models.Alarm, error) {
	return s.repository.FindAll(userId)
}

func (s *AlarmService) FindById(id uint) (models.Alarm, error) {
	alarm, err := s.repository.FindById(id)
	if err != nil {
		return models.Alarm{}, err
	}
	return *alarm, nil
}

func (s *AlarmService) Create(dto dto.CreateAlarmDto) (models.Alarm, error) {
	alarm, err := s.repository.Create(dto)
	if err != nil {
		return models.Alarm{}, err
	}
	return *alarm, nil
}

func (s *AlarmService) Update(dto dto.UpdateAlarmDto) (models.Alarm, error) {
	oldAlarm, err := s.repository.FindById(dto.ID)
	if err != nil {
		return models.Alarm{}, err
	}
	alarm := mappers.UpdateAlarmDtoToAlarmModel(&dto)
	updatedAlarm, err := s.repository.Update(oldAlarm, alarm)
	if err != nil {
		return models.Alarm{}, err
	}
	return *updatedAlarm, nil
}

func (s *AlarmService) Delete(id uint) error {
	return s.repository.Delete(id)
}
