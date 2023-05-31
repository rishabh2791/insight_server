package application

import (
	"insight/pkg/domain/entity"
	"insight/pkg/domain/repository"
)

type DeviceTypeApp struct {
	deviceTypeRepo repository.DeviceTypeRepository
}

func NewDeviceTypeApp() *DeviceTypeApp {
	deviceTypeApp := DeviceTypeApp{}
	return &deviceTypeApp
}

func (app *DeviceTypeApp) Create(device *entity.DeviceType) (*entity.DeviceType, error) {
	return app.deviceTypeRepo.Create(device)
}

func (app *DeviceTypeApp) Get(id string) (*entity.DeviceType, error) {
	return app.deviceTypeRepo.Get(id)
}

func (app *DeviceTypeApp) List(conditions string) ([]entity.DeviceType, error) {
	return app.deviceTypeRepo.List(conditions)
}

func (app *DeviceTypeApp) Update(id string, deviceType *entity.DeviceType) (*entity.DeviceType, error) {
	return app.deviceTypeRepo.Update(id, deviceType)
}

type DeviceTypeAppInterface interface {
	Create(*entity.DeviceType) (*entity.DeviceType, error)
	Get(string) (*entity.DeviceType, error)
	List(string) ([]entity.DeviceType, error)
	Update(string, *entity.DeviceType) (*entity.DeviceType, error)
}
