package persistance

import (
	"insight/pkg/domain/entity"
	"insight/pkg/domain/repository"

	"gorm.io/gorm"
)

type VesselRepo struct {
	DB *gorm.DB
}

func NewVesselRepo(db *gorm.DB) repository.VesselRepository {
	repo := VesselRepo{}
	repo.DB = db
	return &repo
}

func (repo *VesselRepo) Create(vessel *entity.Vessel) (*entity.Vessel, error) {
	return vessel, nil
}

func (repo *VesselRepo) Get(id string) (*entity.Vessel, error) {
	vessel := entity.Vessel{}

	return &vessel, nil
}

func (repo *VesselRepo) List(conditions string) ([]entity.Vessel, error) {
	vessels := []entity.Vessel{}

	return vessels, nil
}

func (repo *VesselRepo) Update(id string, vessel *entity.Vessel) (*entity.Vessel, error) {
	return vessel, nil
}
