package entity

type Device struct {
	ID            string      `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	DeviceTypeID  string      `json:"device_type_id" gorm:"size:191;not null;"`
	DeviceType    *DeviceType `json:"device_type"`
	VesselID      string      `json:"vessel_id" gorm:"size:191;not null;"`
	Vessel        *Vessel     `json:"vessel"`
	IsConstant    bool        `json:"is_constant" gorm:"default:false;"`
	ConstantValue float32     `json:"constant_value"`
	NodeAddress   int         `json:"node_address"`
	ReadStart     int         `json:"read_start"`
	BaudRate      int         `json:"baud_rate"`
	ByteSize      int         `json:"byte_size"`
	StopBits      int         `json:"stop_bits"`
	TimeOut       float32     `json:"time_out"`
	MessageLength int         `json:"message_length"`
	ClearBuffer   bool        `json:"clear_buffers_before_each_transaction" gorm:"default:True;"`
	ClosePort     bool        `json:"close_port_after_each_call" gorm:"default:True;"`
}

func (Device) Tablename() string {
	return "devices"
}
