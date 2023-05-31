package entity

import "time"

type DeviceData struct {
	ID        string    `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	DeviceID  string    `json:"device_id" gorm:"size:191;not null;"`
	Value     float32   `json:"value" gorm:"default:0.0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (DeviceData) Tablename() string {
	return "device_data"
}
