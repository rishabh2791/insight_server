package application

import (
	"insight/pkg/domain/entity"
	"insight/pkg/domain/repository"
)

type VesselApp struct {
	vesselRepo repository.VesselRepository
}

func NewVesselApp() *VesselApp {
	vesselApp := VesselApp{}
	return &vesselApp
}

func (app *VesselApp) Create(vessel *entity.Vessel) (*entity.Vessel, error) {
	return app.vesselRepo.Create(vessel)
}

func (app *VesselApp) Get(id string) (*entity.Vessel, error) {
	return app.vesselRepo.Get(id)
}

func (app *VesselApp) List(conditions string) ([]entity.Vessel, error) {
	return app.vesselRepo.List(conditions)
}

func (app *VesselApp) Update(id string, vessel *entity.Vessel) (*entity.Vessel, error) {
	return app.vesselRepo.Update(id, vessel)
}

type VesselAppInterface interface {
	Create(*entity.Vessel) (*entity.Vessel, error)
	Get(string) (*entity.Vessel, error)
	List(string) ([]entity.Vessel, error)
	Update(string, *entity.Vessel) (*entity.Vessel, error)
}
