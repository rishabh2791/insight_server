package views

import "insight/pkg/application"

type VesselViews struct {
	Application application.Application
}

func NewVesselApp() *VesselViews {
	vesselViews := VesselViews{}

	return &vesselViews
}
