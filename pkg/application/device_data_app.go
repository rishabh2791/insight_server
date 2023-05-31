package application

import (
	"insight/pkg/domain/entity"
	"insight/pkg/domain/repository"
)

type DeviceDataApp struct {
	deviceDataRepo repository.DeviceDataRepository
}

func NewDeviceDataApp() *DeviceDataApp {
	deviceDataApp := DeviceDataApp{}
	return &deviceDataApp
}

func (app *DeviceDataApp) Create(device *entity.DeviceData) (*entity.DeviceData, error) {
	return app.deviceDataRepo.Create(device)
}

func (app *DeviceDataApp) Get(id string) (*entity.DeviceData, error) {
	return app.deviceDataRepo.Get(id)
}

func (app *DeviceDataApp) List(conditions string) ([]entity.DeviceData, error) {
	return app.deviceDataRepo.List(conditions)
}

type DeviceDataAppInterface interface {
	Create(*entity.DeviceData) (*entity.DeviceData, error)
	Get(string) (*entity.DeviceData, error)
	List(string) ([]entity.DeviceData, error)
}
