package server

import (
	"insight/pkg/views"

	"github.com/gin-gonic/gin"
)

type VesselRouter struct {
	router *gin.RouterGroup
	views  *views.ViewStore
}

func NewVesselRouter(routerGroup *gin.RouterGroup, views *views.ViewStore) *VesselRouter {
	router := VesselRouter{}

	router.router = routerGroup
	router.views = views

	return &router
}

func (router *VesselRouter) ServeRoutes() {
	router.router.GET("/test/", router.views.VesselViews.Test)
}
