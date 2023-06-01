package entity

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vessel struct {
	ID   string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Name string `json:"name" gorm:"size:200;not null;unique;"`
}

func (Vessel) Tablename() string {
	return "vessels"
}

func (model *Vessel) BeforeCreate(db *gorm.DB) error {
	model.ID = uuid.New().String()
	return nil
}

func (model *Vessel) Validate() error {
	errs := ""

	if model.Name == "" || len(model.Name) == 0 {
		errs += "Vessel Name Can Not be Empty.\n"
	}

	if len(errs) == 0 || errs == "" {
		return nil
	}

	return errors.New(errs)
}
