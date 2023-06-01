package server

import (
	"insight/pkg/views"

	"github.com/gin-gonic/gin"
)

type DeviceTypeRouter struct {
	router *gin.RouterGroup
	views  *views.ViewStore
}

func NewDeviceTypeRouter(routerGroup *gin.RouterGroup, views *views.ViewStore) *DeviceTypeRouter {
	router := DeviceTypeRouter{}

	router.router = routerGroup
	router.views = views

	return &router
}

func (router *DeviceTypeRouter) ServeRoutes() {
	router.router.POST("/create/", router.views.DeviceTypeViews.Create)
	router.router.GET("/:id/", router.views.DeviceTypeViews.Get)
	router.router.POST("/", router.views.DeviceTypeViews.List)
	router.router.PATCH("/:id/", router.views.DeviceTypeViews.Update)
}
