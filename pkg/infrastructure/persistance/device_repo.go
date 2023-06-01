package persistance

import (
	"insight/pkg/domain/entity"
	"insight/pkg/domain/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	validationErr := device.Validate()

	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := repo.DB.Create(&device).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return device, nil
}

func (repo *DeviceRepo) Get(id string) (*entity.Device, error) {
	device := entity.Device{}

	getErr := repo.DB.Where("id = ?", id).Take(&device).Error

	return &device, getErr
}

func (repo *DeviceRepo) List(conditions string) ([]entity.Device, error) {
	devices := []entity.Device{}

	getErr := repo.DB.
		Preload(clause.Associations).
		Where(conditions).Find(&devices).Error

	return devices, getErr
}

func (repo *DeviceRepo) Update(id string, device *entity.Device) (*entity.Device, error) {
	existingDevice := entity.Device{}

	getErr := repo.DB.Where("id = ?", id).Take(&existingDevice).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := repo.DB.Table(device.Tablename()).Where("id = ?", id).Updates(&device).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.Device{}
	repo.DB.Where("id = ?", id).Take(&updated)

	return &updated, nil
}
