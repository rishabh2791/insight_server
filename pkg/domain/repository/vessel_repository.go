package repository

import "insight/pkg/domain/entity"

type VesselRepository interface {
	Create(*entity.Vessel) (*entity.Vessel, error)
	Get(string) (*entity.Vessel, error)
	List(string) ([]entity.Vessel, error)
	Update(string, *entity.Vessel) (*entity.Vessel, error)
}
