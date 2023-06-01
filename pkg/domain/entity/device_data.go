package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceData struct {
	ID        string    `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	DeviceID  string    `json:"device_id" gorm:"size:191;not null;"`
	Device    *Device   `json:"device"`
	Value     float32   `json:"value" gorm:"default:0.0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (DeviceData) Tablename() string {
	return "device_data"
}

func (model *DeviceData) BeforeCreate(db *gorm.DB) error {
	model.ID = uuid.New().String()
	return nil
}

func (model *DeviceData) Validate() error {
	errs := ""

	if model.DeviceID == "" || len(model.DeviceID) == 0 {
		errs += "Device Can Not be Empty.\n"
	}

	if len(errs) == 0 || errs == "" {
		return nil
	}

	return errors.New(errs)
}
