package server

import (
	"insight/pkg/views"

	"github.com/gin-gonic/gin"
)

type DeviceDataRouter struct {
	router *gin.RouterGroup
	views  *views.ViewStore
}

func NewDeviceDataRouter(routerGroup *gin.RouterGroup, views *views.ViewStore) *DeviceDataRouter {
	router := DeviceDataRouter{}

	router.router = routerGroup
	router.views = views

	return &router
}

func (router *DeviceDataRouter) ServeRoutes() {
	router.router.POST("/create/", router.views.DeviceDataViews.Create)
	router.router.GET("/:id/", router.views.DeviceDataViews.Get)
	router.router.POST("/", router.views.DeviceDataViews.List)
}
