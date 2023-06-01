package application

import "insight/pkg/infrastructure/persistance"

type Application struct {
	VesselApp     *VesselApp
	DeviceTypeApp *DeviceTypeApp
	DeviceApp     *DeviceApp
	DeviceDataApp *DeviceDataApp
}

func NewApplication(repoStore persistance.RepoStore) *Application {
	app := Application{}
	app.VesselApp = NewVesselApp(repoStore.VesselRepo)
	app.DeviceApp = NewDeviceApp(repoStore.DeviceRepo)
	app.DeviceTypeApp = NewDeviceTypeApp(repoStore.DeviceTypeRepo)
	app.DeviceDataApp = NewDeviceDataApp(repoStore.DeviceDataRepo)
	return &app
}
