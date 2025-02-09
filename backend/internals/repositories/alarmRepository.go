package repositories

import (
	"github.com/iacopoghilardi/alarm-tool/internals/mappers"
	"github.com/iacopoghilardi/alarm-tool/internals/models"
	"github.com/iacopoghilardi/alarm-tool/internals/types/dto"
	"gorm.io/gorm"
)

type AlarmRepository struct {
	db *gorm.DB
}

func NewAlarmRepository(db *gorm.DB) *AlarmRepository {
	return &AlarmRepository{db: db}
}

func (r *AlarmRepository) FindAll(userId uint) ([]models.Alarm, error) {
	var alarms []models.Alarm
	if err := r.db.Where("user_id = ?", userId).Find(&alarms).Error; err != nil {
		return nil, err
	}
	return alarms, nil
}

func (r *AlarmRepository) FindById(id uint) (*models.Alarm, error) {
	var alarm models.Alarm
	if err := r.db.Where("id = ?", id).First(&alarm).Error; err != nil {
		return nil, err
	}
	return &alarm, nil
}

func (r *AlarmRepository) Create(dto dto.CreateAlarmDto) (*models.Alarm, error) {
	alarm := mappers.CreateAlarmDtoToAlarmModel(&dto)

	if err := r.db.Create(&alarm).Error; err != nil {
		return nil, err
	}
	return &alarm, nil
}

func (r *AlarmRepository) Update(oldAlarm *models.Alarm, alarm *models.Alarm) (*models.Alarm, error) {
	if err := r.db.Model(&oldAlarm).Updates(alarm).Error; err != nil {
		return nil, err
	}
	return oldAlarm, nil
}

func (r *AlarmRepository) Delete(id uint) error {
	return r.db.Delete(&models.Alarm{}, id).Error
}
