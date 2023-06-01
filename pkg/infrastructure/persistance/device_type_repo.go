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
	validationErr := deviceType.Validate()

	if validationErr != nil {
		return nil, validationErr
	}

	creationErr := repo.DB.Create(&deviceType).Error
	if creationErr != nil {
		return nil, creationErr
	}

	return deviceType, nil
}

func (repo *DeviceTypeRepo) Get(id string) (*entity.DeviceType, error) {
	deviceType := entity.DeviceType{}

	getErr := repo.DB.Where("id = ?", id).Take(&deviceType).Error

	return &deviceType, getErr
}

func (repo *DeviceTypeRepo) List(conditions string) ([]entity.DeviceType, error) {
	deviceTypes := []entity.DeviceType{}

	getErr := repo.DB.Where(conditions).Find(&deviceTypes).Error

	return deviceTypes, getErr
}

func (repo *DeviceTypeRepo) Update(id string, deviceType *entity.DeviceType) (*entity.DeviceType, error) {
	existingDeviceType := entity.DeviceType{}

	getErr := repo.DB.Where("id = ?", id).Take(&existingDeviceType).Error
	if getErr != nil {
		return nil, getErr
	}

	updationErr := repo.DB.Table(deviceType.Tablename()).Where("id = ?", id).Updates(&deviceType).Error
	if updationErr != nil {
		return nil, updationErr
	}

	updated := entity.DeviceType{}
	repo.DB.Where("id = ?", id).Take(&updated)

	return &updated, nil
}
