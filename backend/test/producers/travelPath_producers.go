package producers

import (
	"log"

	"github.com/iacopoghilardi/alarm-tool/internals/models"
	"gorm.io/gorm"
)

func DeleteAllTravelPaths(db *gorm.DB) {
	db.Exec("DELETE FROM travel_paths")
}

func DeleteTestTravelPath(db *gorm.DB) {
	db.Where("label = ?", "Test Travel Path").Delete(&models.TravelPath{})
}

func CreateTestTravelPath(db *gorm.DB) *models.TravelPath {
	travelPath := &models.TravelPath{
		StartLocation: models.Location{
			Lat: 1.0,
			Lng: 1.0,
		},
		EndLocation: models.Location{
			Lat: 2.0,
			Lng: 2.0,
		},
	}

	if err := db.Create(travelPath).Error; err != nil {
		log.Fatalf("Failed to create travel path: %v", err)
	}

	return travelPath
}
