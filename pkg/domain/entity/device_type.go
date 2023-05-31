package entity

/*
	Options:
	1.	Anchor Speed
	2.	Cowl Speed
	3.	Paddle Speed
	4.	Inner Speed
	5.	Emulsifier Speed
	6.	Main Vessel Temperature
	7.	Hot Pot Temperature
	8.	Main Vessel Pressure
	9.	Main Vessel Load Cell
	10. Hot Pot Load Cell
*/

type DeviceType struct {
	ID          string `json:"id" gorm:"size:191;primaryKey;unique;not null;"`
	Description string `json:"description" gorm:"size:200;not null;"`
}

func (DeviceType) Tablename() string {
	return "device_types"
}
