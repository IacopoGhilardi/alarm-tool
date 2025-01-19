package producers

import (
	"log"

	"github.com/iacopoghilardi/alarm-tool/internals/models"
	"gorm.io/gorm"
)

func DeleteAllAlarms(db *gorm.DB) {
	db.Exec("DELETE FROM alarms")
}

func DeleteTestAlarm(db *gorm.DB) {
	db.Where("label = ?", "Test Alarm").Delete(&models.Alarm{})
}

func CreateTestAlarm(db *gorm.DB) *models.Alarm {
	alarm := &models.Alarm{
		Label: "Test Alarm",
		Time:  "10:00",
	}

	if err := db.Create(alarm).Error; err != nil {
		log.Fatalf("Failed to create alarm: %v", err)
	}

	return alarm
}
