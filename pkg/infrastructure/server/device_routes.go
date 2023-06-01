package server

import (
	"insight/pkg/views"

	"github.com/gin-gonic/gin"
)

type DeviceRouter struct {
	router *gin.RouterGroup
	views  *views.ViewStore
}

func NewDeviceRouter(routerGroup *gin.RouterGroup, views *views.ViewStore) *DeviceRouter {
	router := DeviceRouter{}

	router.router = routerGroup
	router.views = views

	return &router
}

func (router *DeviceRouter) ServeRoutes() {
	router.router.POST("/create/", router.views.DeviceViews.Create)
	router.router.GET("/:id/", router.views.DeviceViews.Get)
	router.router.POST("/", router.views.DeviceViews.List)
	router.router.PATCH("/:id/", router.views.DeviceViews.Update)
}
