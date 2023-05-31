package entity

type Vessel struct {
	ID   string `json:"id" gorm:"size:191;not null;primaryKey;unique;"`
	Name string `json:"vessel_name" gorm:"size:200;not null;unique;"`
}

func (Vessel) Tablename() string {
	return "vessels"
}
