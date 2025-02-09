package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/iacopoghilardi/alarm-tool/internals/repositories"
	"github.com/iacopoghilardi/alarm-tool/internals/services"
	"github.com/iacopoghilardi/alarm-tool/internals/types/dto"
	"github.com/iacopoghilardi/alarm-tool/test/producers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"gorm.io/gorm"
)

type AlarmTestSuite struct {
	suite.Suite
	db          *gorm.DB
	pgContainer *postgres.PostgresContainer
	service     *services.AlarmService
	ctx         context.Context
}

func (s *AlarmTestSuite) SetupSuite() {
	s.ctx = context.Background()

	pgContainer, db, err := producers.SetupGenericSuite(s.ctx)
	if err != nil {
		s.T().Fatalf("Failed to setup generic suite: %v", err)
	}
	s.pgContainer = pgContainer
	s.db = db
	s.service = services.NewAlarmService(repositories.NewAlarmRepository(s.db))
}

func (s *AlarmTestSuite) TearDownSuite() {
	producers.TearDownGenericSuite(s.ctx, s.db, s.pgContainer)
}

func (s *AlarmTestSuite) TestCreateAlarm() {
	producers.DeleteAllAlarms(s.db)

	createAlarmDto := dto.CreateAlarmDto{
		Label:           "Test Alarm",
		Time:            "10:00",
		IsActive:        true,
		IsRepeat:        false,
		RepeatDays:      []string{},
		RepeatType:      "daily",
		RepeatEndTime:   time.Now().Add(24 * time.Hour),
		RepeatStartTime: time.Now(),
		SnoozeEnabled:   false,
		SnoozeInterval:  5,
		SmartWakeup:     false,
		WakeupWindow:    10,
		WeatherAware:    false,
	}

	alarm, err := s.service.Create(createAlarmDto)
	if err != nil {
		log.Fatal("Failed to create alarm: ", err)
	}

	assert.NotNil(s.T(), alarm)
	assert.Equal(s.T(), alarm.Label, createAlarmDto.Label)
	assert.Equal(s.T(), alarm.Time, createAlarmDto.Time)
	assert.Equal(s.T(), alarm.IsActive, createAlarmDto.IsActive)
	assert.Equal(s.T(), alarm.IsRepeat, createAlarmDto.IsRepeat)
	assert.Equal(s.T(), alarm.RepeatDays, createAlarmDto.RepeatDays)
	assert.Equal(s.T(), alarm.RepeatType, createAlarmDto.RepeatType)
	assert.Equal(s.T(), alarm.RepeatEndTime, createAlarmDto.RepeatEndTime)
	assert.Equal(s.T(), alarm.RepeatStartTime, createAlarmDto.RepeatStartTime)
	assert.Equal(s.T(), alarm.SnoozeEnabled, createAlarmDto.SnoozeEnabled)
	assert.Equal(s.T(), alarm.SnoozeInterval, createAlarmDto.SnoozeInterval)
	assert.Equal(s.T(), alarm.SmartWakeup, createAlarmDto.SmartWakeup)
	assert.Equal(s.T(), alarm.WakeupWindow, createAlarmDto.WakeupWindow)
	assert.Equal(s.T(), alarm.WeatherAware, createAlarmDto.WeatherAware)
}

func (s *AlarmTestSuite) TestUpdateAlarm() {
	producers.DeleteAllAlarms(s.db)

	alarm := producers.CreateTestAlarm(s.db)

	updateAlarmDto := dto.UpdateAlarmDto{
		ID:              alarm.ID,
		Label:           "Updated Alarm",
		Time:            "11:00",
		IsActive:        true,
		IsRepeat:        false,
		RepeatDays:      []string{},
		RepeatType:      "daily",
		RepeatEndTime:   time.Now().Add(24 * time.Hour),
		RepeatStartTime: time.Now(),
	}

	updatedAlarm, err := s.service.Update(updateAlarmDto)
	if err != nil {
		log.Fatal("Failed to update alarm: ", err)
	}

	assert.NotNil(s.T(), updatedAlarm)
	assert.Equal(s.T(), updatedAlarm.Label, updateAlarmDto.Label)
	assert.Equal(s.T(), updatedAlarm.Time, updateAlarmDto.Time)
	assert.Equal(s.T(), updatedAlarm.IsActive, updateAlarmDto.IsActive)
	assert.Equal(s.T(), updatedAlarm.IsRepeat, updateAlarmDto.IsRepeat)
	assert.Equal(s.T(), updatedAlarm.RepeatDays, updateAlarmDto.RepeatDays)
	assert.Equal(s.T(), updatedAlarm.RepeatType, updateAlarmDto.RepeatType)
	assert.Equal(s.T(), updatedAlarm.RepeatEndTime, updateAlarmDto.RepeatEndTime)
	assert.Equal(s.T(), updatedAlarm.RepeatStartTime, updateAlarmDto.RepeatStartTime)
}

func (s *AlarmTestSuite) TestDeleteAlarm() {
	producers.DeleteAllAlarms(s.db)

	alarm := producers.CreateTestAlarm(s.db)

	err := s.service.Delete(alarm.ID)
	if err != nil {
		log.Fatal("Failed to delete alarm: ", err)
	}

	deletedAlarm, err := s.service.FindById(alarm.ID)
	assert.Nil(s.T(), deletedAlarm)
	assert.Error(s.T(), err)
}

func TestAlarmTestSuite(t *testing.T) {
	suite.Run(t, new(AlarmTestSuite))
}
