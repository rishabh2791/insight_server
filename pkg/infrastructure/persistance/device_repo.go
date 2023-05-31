package persistance

import (
	"insight/pkg/domain/entity"
	"insight/pkg/domain/repository"

	"gorm.io/gorm"
)

type DeviceRepo struct {
	DB *gorm.DB
}

func NewDeviceRepo(db *gorm.DB) repository.DeviceRepository {
	repo := DeviceRepo{}
	repo.DB = db
	return &repo
}

func (repo *DeviceRepo) Create(device *entity.Device) (*entity.Device, error) {
	return device, nil
}

func (repo *DeviceRepo) Get(id string) (*entity.Device, error) {
	device := entity.Device{}

	return &device, nil
}

func (repo *DeviceRepo) List(string) ([]entity.Device, error) {
	devices := []entity.Device{}

	return devices, nil
}

func (repo *DeviceRepo) Update(id string, device *entity.Device) (*entity.Device, error) {
	return device, nil
}
