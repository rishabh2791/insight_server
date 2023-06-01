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
	validationErr := vessel.Validate()

	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := repo.DB.Create(&vessel).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return vessel, nil
}

func (repo *VesselRepo) Get(id string) (*entity.Vessel, error) {
	vessel := entity.Vessel{}

	getErr := repo.DB.Where("id = ?", id).Take(&vessel).Error

	return &vessel, getErr
}

func (repo *VesselRepo) List(conditions string) ([]entity.Vessel, error) {
	vessels := []entity.Vessel{}

	getErr := repo.DB.Where(conditions).Find(&vessels).Error

	return vessels, getErr
}

func (repo *VesselRepo) Update(id string, vessel *entity.Vessel) (*entity.Vessel, error) {
	existingVessel := entity.Vessel{}

	getErr := repo.DB.Where("id = ?", id).Take(&existingVessel).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := repo.DB.Table(vessel.Tablename()).Where("id = ?", id).Updates(&vessel).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Vessel{}
	repo.DB.Where("id = ?", id).Take(&updated)
	return &updated, nil
}
