package server

import (
	"fmt"
	"insight/pkg/views"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router    *gin.Engine
	Hostname  string
	Port      int
	ViewStore *views.ViewStore
}

func NewServer(host string, port int, views *views.ViewStore) *Server {
	server := Server{}

	server.Router = gin.Default()
	server.Hostname = host
	server.Port = port
	server.ViewStore = views

	return &server
}

func (server *Server) Run() *http.Server {
	httpServer := http.Server{
		Addr:         fmt.Sprintf("%s:%v", server.Hostname, server.Port),
		Handler:      server.Router,
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
		IdleTimeout:  300 * time.Second,
	}
	return &httpServer
}

func (server *Server) Serve() {
	vesselRouter := NewVesselRouter(server.Router.Group("/vessel/"), server.ViewStore)
	deviceTypeRouter := NewDeviceTypeRouter(server.Router.Group("/devicetype/"), server.ViewStore)
	deviceRouter := NewDeviceRouter(server.Router.Group("/device/"), server.ViewStore)
	deviceDataRouter := NewDeviceDataRouter(server.Router.Group("/devicedata/"), server.ViewStore)

	vesselRouter.ServeRoutes()
	deviceTypeRouter.ServeRoutes()
	deviceRouter.ServeRoutes()
	deviceDataRouter.ServeRoutes()
}
