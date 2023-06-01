package persistance

import (
	"insight/pkg/domain/entity"
	"insight/pkg/domain/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	validationErr := deviceData.Validate()

	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := repo.DB.Create(&deviceData).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return deviceData, nil
}

func (repo *DeviceDataRepo) Get(id string) (*entity.DeviceData, error) {
	deviceData := entity.DeviceData{}

	getErr := repo.DB.Where("id = ?", id).Take(&deviceData).Error

	return &deviceData, getErr
}

func (repo *DeviceDataRepo) List(conditions string) ([]entity.DeviceData, error) {
	deviceData := []entity.DeviceData{}

	getErr := repo.DB.
		Preload("Device.Vessel").
		Preload("Device.DeviceType").
		Preload(clause.Associations).
		Where(conditions).Find(&deviceData).Error

	return deviceData, getErr
}
