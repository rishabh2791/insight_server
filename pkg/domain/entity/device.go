package entity

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	ID                    string      `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	DeviceTypeID          string      `json:"device_type_id" gorm:"size:191;not null;"`
	DeviceType            *DeviceType `json:"device_type"`
	VesselID              string      `json:"vessel_id" gorm:"size:191;not null;"`
	Vessel                *Vessel     `json:"vessel"`
	Port                  string      `json:"port" gorm:"size:100;not null;default:'/dev/ttyAgitators'"`
	IsConstant            bool        `json:"is_constant" gorm:"default:false;"`
	ConstantValue         float32     `json:"constant_value" gorm:"default:0.0;"`
	Factor                int         `json:"factor" gorm:"default:1;"`
	NodeAddress           int         `json:"node_address" gorm:"default:1;"`
	AdditionalNodeAddress int         `json:"additional_node_address" gorm:"default:22;"` // Can be used as GPIO Pin if isConstant is true.
	ReadStart             int         `json:"read_start" gorm:"default:0;"`
	BaudRate              int         `json:"baud_rate"  gorm:"default:9600;"`
	ByteSize              int         `json:"byte_size"  gorm:"default:16;"`
	StopBits              int         `json:"stop_bits"  gorm:"default:2;"`
	TimeOut               float32     `json:"time_out"  gorm:"default:0.5;"`
	Enabled               bool        `json:"enabled" gorm:"default:true;"`
	MessageLength         int         `json:"message_length" gorm:"default:16;"`
	ClearBuffer           bool        `json:"clear_buffers_before_each_transaction" gorm:"default:True;"`
	ClosePort             bool        `json:"close_port_after_each_call" gorm:"default:True;"`
	CommunicationMethod   string      `json:"communication_method" gorm:"size:20;not null;default:'modbus'"` // Options between modbus, serial and constant (GPIO)
}

func (Device) Tablename() string {
	return "devices"
}

func (model *Device) BeforeCreate(db *gorm.DB) error {
	model.ID = uuid.New().String()
	return nil
}

func (model *Device) Validate() error {
	errs := ""

	if model.DeviceTypeID == "" || len(model.DeviceTypeID) == 0 {
		errs += "Device Type Can Not be Empty.\n"
	}

	if model.VesselID == "" || len(model.VesselID) == 0 {
		errs += "Vessel Can Not be Empty.\n"
	}

	if model.CommunicationMethod == "" || len(model.CommunicationMethod) == 0 {
		errs += "Communication Method Can Not be Empty.\n"
	}

	if len(errs) == 0 || errs == "" {
		return nil
	}

	return errors.New(errs)
}
