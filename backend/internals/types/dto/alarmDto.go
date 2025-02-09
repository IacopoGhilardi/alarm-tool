package dto

import "time"

type AlarmDto struct {
	ID              uint      `json:"id"`
	CreatedAt       string    `json:"createdAt"`
	UpdatedAt       string    `json:"updatedAt"`
	UserID          uint      `json:"user_id"`
	Label           string    `json:"label"`
	Time            string    `json:"time"`
	IsActive        bool      `json:"isActive"`
	IsRepeat        bool      `json:"isRepeat"`
	RepeatDays      []string  `json:"repeat_days"`
	RepeatType      string    `json:"repeat_type"`
	RepeatEndTime   time.Time `json:"repeat_end_time"`
	RepeatStartTime time.Time `json:"repeat_start_time"`
	Sound           string    `json:"sound"`
	Vibration       bool      `json:"vibration"`
	Volume          int       `json:"volume"`
	SnoozeEnabled   bool      `json:"snooze_enabled"`
	SnoozeInterval  int       `json:"snooze_interval"`
	SmartWakeup     bool      `json:"smart_wakeup"`
	WakeupWindow    int       `json:"wakeup_window"`
	WeatherAware    bool      `json:"weather_aware"`
}

type CreateAlarmDto struct {
	UserID          uint      `json:"user_id"`
	Label           string    `json:"label"`
	Time            string    `json:"time"`
	IsActive        bool      `json:"isActive"`
	IsRepeat        bool      `json:"isRepeat"`
	RepeatDays      []string  `json:"repeat_days"`
	RepeatType      string    `json:"repeat_type"`
	RepeatEndTime   time.Time `json:"repeat_end_time"`
	RepeatStartTime time.Time `json:"repeat_start_time"`
	SnoozeEnabled   bool      `json:"snooze_enabled"`
	SnoozeInterval  int       `json:"snooze_interval"`
	SmartWakeup     bool      `json:"smart_wakeup"`
	WakeupWindow    int       `json:"wakeup_window"`
	WeatherAware    bool      `json:"weather_aware"`
}

type UpdateAlarmDto struct {
	ID              uint      `json:"id"`
	UpdatedAt       string    `json:"updatedAt"`
	UserID          uint      `json:"user_id"`
	Label           string    `json:"label"`
	Time            string    `json:"time"`
	IsActive        bool      `json:"isActive"`
	IsRepeat        bool      `json:"isRepeat"`
	RepeatDays      []string  `json:"repeat_days"`
	RepeatType      string    `json:"repeat_type"`
	RepeatEndTime   time.Time `json:"repeat_end_time"`
	RepeatStartTime time.Time `json:"repeat_start_time"`
	Sound           string    `json:"sound"`
	Vibration       bool      `json:"vibration"`
	Volume          int       `json:"volume"`
	SnoozeEnabled   bool      `json:"snooze_enabled"`
	SnoozeInterval  int       `json:"snooze_interval"`
	SmartWakeup     bool      `json:"smart_wakeup"`
	WakeupWindow    int       `json:"wakeup_window"`
	WeatherAware    bool      `json:"weather_aware"`
}
