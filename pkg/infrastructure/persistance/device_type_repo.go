package persistance

import (
	"insight/pkg/domain/entity"
	"insight/pkg/domain/repository"

	"gorm.io/gorm"
)

type DeviceTypeRepo struct {
	DB *gorm.DB
}

func NewDeviceTypeRepo(db *gorm.DB) repository.DeviceTypeRepository {
	repo := DeviceTypeRepo{}
	repo.DB = db
	return &repo
}

func (repo *DeviceTypeRepo) Create(deviceType *entity.DeviceType) (*entity.DeviceType, error) {
	return deviceType, nil
}

func (repo *DeviceTypeRepo) Get(id string) (*entity.DeviceType, error) {
	deviceType := entity.DeviceType{}

	return &deviceType, nil
}

func (repo *DeviceTypeRepo) List(string) ([]entity.DeviceType, error) {
	deviceTypes := []entity.DeviceType{}

	return deviceTypes, nil
}

func (repo *DeviceTypeRepo) Update(id string, deviceType *entity.DeviceType) (*entity.DeviceType, error) {
	return deviceType, nil
}
