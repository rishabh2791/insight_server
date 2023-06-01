package views

import "insight/pkg/application"

type DeviceDataViews struct {
	apps *application.Application
}

func NewDeviceDataViews(apps *application.Application) *DeviceDataViews {
	views := DeviceDataViews{}

	views.apps = apps

	return &views
}
