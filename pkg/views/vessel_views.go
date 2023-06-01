package views

import (
	"insight/pkg/application"
	"log"

	"github.com/gin-gonic/gin"
)

type VesselViews struct {
	apps *application.Application
}

func NewVesselViews(apps *application.Application) *VesselViews {
	views := VesselViews{}

	views.apps = apps

	return &views
}

func (view *VesselViews) Test(ctx *gin.Context) {
	log.Println("Testing")
}
