package models

import "time"

type TravelPath struct {
	ID            uint      `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	UserID        uint      `gorm:"column:user_id" json:"user_id"`
	User          User      `gorm:"foreignKey:UserID" json:"user"`
	StartLocation Location  `gorm:"column:start_location" json:"start_location"`
	EndLocation   Location  `gorm:"column:end_location" json:"end_location"`
}

type Location struct {
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Address string  `json:"address"`
}
