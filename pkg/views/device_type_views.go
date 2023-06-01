package views

import "insight/pkg/application"

type DeviceTypeViews struct {
	apps *application.Application
}

func NewDeviceTypeViews(apps *application.Application) *DeviceTypeViews {
	views := DeviceTypeViews{}

	views.apps = apps

	return &views
}
