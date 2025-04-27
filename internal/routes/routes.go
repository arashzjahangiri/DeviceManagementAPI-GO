package routes

import (
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/controllers"
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Router manages the application routes
type Router struct {
	engine     *gin.Engine
	log        *logrus.Logger
	deviceCtrl *controllers.DevicesController
}

// NewRouter creates a new Router instance
func NewRouter(deviceService *services.DeviceService, log *logrus.Logger) *Router {
	engine := gin.Default()
	deviceCtrl := controllers.NewDevicesController(deviceService, log)

	return &Router{
		engine:     engine,
		log:        log,
		deviceCtrl: deviceCtrl,
	}
}

// SetupRoutes registers all application routes
func (r *Router) SetupRoutes() {
	r.log.Info("Setting up routes")
	api := r.engine.Group("/api/v1")
	{
		devices := api.Group("/devices")
		{
			devices.GET("", r.deviceCtrl.GetDevices)
			devices.GET(":id", r.deviceCtrl.GetDevice)
			devices.POST("", r.deviceCtrl.CreateDevice)
			devices.PUT(":id", r.deviceCtrl.UpdateDevice)
			devices.DELETE(":id", r.deviceCtrl.DeleteDevice)
		}
	}
}

// Run starts the HTTP server
func (r *Router) Run(addr string) error {
	r.log.Infof("Starting server on %s", addr)
	return r.engine.Run(addr)
}
