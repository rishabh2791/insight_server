package views

import "insight/pkg/application"

type ViewStore struct {
	application     *application.Application
	VesselViews     *VesselViews
	DeviceTypeViews *DeviceTypeViews
	DeviceViews     *DeviceViews
	DeviceDataViews *DeviceDataViews
}

func NewViewStore(app *application.Application) *ViewStore {
	viewStore := ViewStore{}

	viewStore.application = app
	viewStore.VesselViews = NewVesselViews(app)
	viewStore.DeviceTypeViews = NewDeviceTypeViews(app)
	viewStore.DeviceViews = NewDeviceViews(app)
	viewStore.DeviceDataViews = NewDeviceDataViews(app)

	return &viewStore
}
