package repositories

import (
	"github.com/iacopoghilardi/alarm-tool/internals/models"
	"gorm.io/gorm"
)

type TravelPathRepository struct {
	db *gorm.DB
}

func NewTravelPathRepository(db *gorm.DB) *TravelPathRepository {
	return &TravelPathRepository{db: db}
}

func (r *TravelPathRepository) FindAll(userId string) ([]models.TravelPath, error) {
	var travelPaths []models.TravelPath
	if err := r.db.Where("user_id = ?", userId).Find(&travelPaths).Error; err != nil {
		return nil, err
	}
	return travelPaths, nil
}

func (r *TravelPathRepository) Create(travelPath *models.TravelPath) (*models.TravelPath, error) {
	if err := r.db.Create(travelPath).Error; err != nil {
		return nil, err
	}
	return travelPath, nil
}

func (r *TravelPathRepository) Delete(id string) error {
	return r.db.Delete(&models.TravelPath{}, id).Error
}

func (r *TravelPathRepository) Update(id string, travelPath *models.TravelPath) (*models.TravelPath, error) {
	if err := r.db.Model(&models.TravelPath{}).Where("id = ?", id).Updates(travelPath).Error; err != nil {
		return nil, err
	}
	var updatedTravelPath models.TravelPath
	if err := r.db.First(&updatedTravelPath, id).Error; err != nil {
		return nil, err
	}
	return &updatedTravelPath, nil
}
