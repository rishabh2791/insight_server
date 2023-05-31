package application

type Application struct {
	VesselApp     *VesselApp
	DeviceTypeApp *DeviceDataApp
	DeviceApp     *DeviceApp
	DeviceDataApp *DeviceDataApp
}

func NewApplication() *Application {
	app := Application{}
	app.VesselApp = NewVesselApp()
	app.DeviceApp = NewDeviceApp()
	app.DeviceTypeApp = NewDeviceDataApp()
	app.DeviceDataApp = NewDeviceDataApp()
	return &app
}
