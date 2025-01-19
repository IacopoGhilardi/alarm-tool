package models

import "time"

type Alarm struct {
	ID        string    `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`

	// Campi base sveglia
	Label    string `gorm:"column:label" json:"label"`         // Nome/etichetta sveglia
	Time     string `gorm:"column:time" json:"time"`           // Orario sveglia (HH:MM)
	IsActive bool   `gorm:"column:is_active" json:"is_active"` // Sveglia attiva/disattiva

	// Gestione ripetizioni
	IsRepeat        bool      `gorm:"column:is_repeat" json:"is_repeat"`             // Sveglia ricorrente
	RepeatDays      []string  `gorm:"column:repeat_days" json:"repeat_days"`         // Giorni di ripetizione
	RepeatType      string    `gorm:"column:repeat_type" json:"repeat_type"`         // Tipo ripetizione (daily, weekly, custom)
	RepeatEndTime   time.Time `gorm:"column:repeat_end_time" json:"repeat_end_time"` // Data fine ripetizione
	RepeatStartTime time.Time `gorm:"column:repeat_start_time" json:"repeat_start_time"`

	// Personalizzazione
	Sound          string `gorm:"column:sound" json:"sound"`                     // Suoneria
	Vibration      bool   `gorm:"column:vibration" json:"vibration"`             // Vibrazione on/off
	Volume         int    `gorm:"column:volume" json:"volume"`                   // Volume sveglia
	SnoozeEnabled  bool   `gorm:"column:snooze_enabled" json:"snooze_enabled"`   // Snooze abilitato
	SnoozeInterval int    `gorm:"column:snooze_interval" json:"snooze_interval"` // Intervallo snooze (minuti)

	// Funzionalità smart
	SmartWakeup  bool `gorm:"column:smart_wakeup" json:"smart_wakeup"`   // Sveglia intelligente
	WakeupWindow int  `gorm:"column:wakeup_window" json:"wakeup_window"` // Finestra di tempo per sveglia smart (minuti)
	WeatherAware bool `gorm:"column:weather_aware" json:"weather_aware"` // Considera condizioni meteo
	TrafficAware bool `gorm:"column:traffic_aware" json:"traffic_aware"` // Considera condizioni traffico
}
