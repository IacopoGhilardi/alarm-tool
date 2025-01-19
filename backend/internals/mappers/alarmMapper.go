package mappers

import (
	"github.com/iacopoghilardi/alarm-tool/internals/models"
	"github.com/iacopoghilardi/alarm-tool/internals/types/dto"
)

func CreateAlarmDtoToAlarmModel(dto *dto.CreateAlarmDto) models.Alarm {
	return models.Alarm{
		Label:           dto.Label,
		Time:            dto.Time,
		UserID:          dto.UserID,
		RepeatDays:      dto.RepeatDays,
		RepeatType:      dto.RepeatType,
		RepeatEndTime:   dto.RepeatEndTime,
		RepeatStartTime: dto.RepeatStartTime,
		SnoozeEnabled:   dto.SnoozeEnabled,
		SnoozeInterval:  dto.SnoozeInterval,
		SmartWakeup:     dto.SmartWakeup,
		WeatherAware:    dto.WeatherAware,
		IsActive:        dto.IsActive,
		IsRepeat:        dto.IsRepeat,
		WakeupWindow:    dto.WakeupWindow,
	}
}

func UpdateAlarmDtoToAlarmModel(dto *dto.UpdateAlarmDto) *models.Alarm {
	return &models.Alarm{
		ID:              dto.ID,
		Label:           dto.Label,
		Time:            dto.Time,
		RepeatDays:      dto.RepeatDays,
		RepeatType:      dto.RepeatType,
		RepeatEndTime:   dto.RepeatEndTime,
		RepeatStartTime: dto.RepeatStartTime,
		SnoozeEnabled:   dto.SnoozeEnabled,
		SnoozeInterval:  dto.SnoozeInterval,
		SmartWakeup:     dto.SmartWakeup,
		WeatherAware:    dto.WeatherAware,
		IsActive:        dto.IsActive,
		IsRepeat:        dto.IsRepeat,
		WakeupWindow:    dto.WakeupWindow,
	}
}
