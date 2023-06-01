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
	router.router.POST("/create/", router.views.VesselViews.Create)
	router.router.GET("/:id/", router.views.VesselViews.Get)
	router.router.POST("/", router.views.VesselViews.List)
	router.router.PATCH("/:id/", router.views.VesselViews.Update)

}
