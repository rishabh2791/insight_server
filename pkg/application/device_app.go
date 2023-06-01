package application

import (
	"insight/pkg/domain/entity"
	"insight/pkg/domain/repository"
)

type DeviceApp struct {
	deviceRepo repository.DeviceRepository
}

func NewDeviceApp(repo repository.DeviceRepository) *DeviceApp {
	deviceApp := DeviceApp{}

	deviceApp.deviceRepo = repo

	return &deviceApp
}

func (app *DeviceApp) Create(device *entity.Device) (*entity.Device, error) {
	return app.deviceRepo.Create(device)
}

func (app *DeviceApp) Get(id string) (*entity.Device, error) {
	return app.deviceRepo.Get(id)
}

func (app *DeviceApp) List(conditions string) ([]entity.Device, error) {
	return app.deviceRepo.List(conditions)
}

func (app *DeviceApp) Update(id string, device *entity.Device) (*entity.Device, error) {
	return app.deviceRepo.Update(id, device)
}

type DeviceAppInterface interface {
	Create(*entity.Device) (*entity.Device, error)
	Get(string) (*entity.Device, error)
	List(string) ([]entity.Device, error)
	Update(string, *entity.Device) (*entity.Device, error)
}
