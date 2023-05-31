package persistance

import (
	"insight/pkg/domain/entity"
	"insight/pkg/domain/repository"

	"gorm.io/gorm"
)

type DeviceDataRepo struct {
	DB *gorm.DB
}

func NewDeviceDataRepo(db *gorm.DB) repository.DeviceDataRepository {
	repo := DeviceDataRepo{}
	repo.DB = db
	return &repo
}

func (repo *DeviceDataRepo) Create(deviceData *entity.DeviceData) (*entity.DeviceData, error) {
	return deviceData, nil
}

func (repo *DeviceDataRepo) Get(id string) (*entity.DeviceData, error) {
	deviceData := entity.DeviceData{}

	return &deviceData, nil
}

func (repo *DeviceDataRepo) List(string) ([]entity.DeviceData, error) {
	deviceDatas := []entity.DeviceData{}

	return deviceDatas, nil
}
