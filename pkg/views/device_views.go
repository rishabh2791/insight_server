package views

import "insight/pkg/application"

type DeviceViews struct {
	apps *application.Application
}

func NewDeviceViews(apps *application.Application) *DeviceViews {
	views := DeviceViews{}

	views.apps = apps

	return &views
}
