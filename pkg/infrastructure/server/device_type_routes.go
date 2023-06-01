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

func (router *DeviceTypeRouter) ServeRoutes() {}
