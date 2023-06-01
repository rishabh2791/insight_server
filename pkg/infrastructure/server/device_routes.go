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

func (router *DeviceRouter) ServeRoutes() {}
