package main

import (
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/config"
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/data"
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/routes"
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/services"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.WithError(err).Fatal("Failed to load configuration")
	}

	// Initialize database
	db, err := data.NewDatabase(cfg, logger)
	if err != nil {
		logger.WithError(err).Fatal("Failed to initialize database")
	}

	// Initialize services
	deviceService := services.NewDeviceService(db, logger)

	// Set up routes
	router := routes.NewRouter(deviceService, logger)
	router.SetupRoutes()

	// Start the server
	if err := router.Run(":8080"); err != nil {
		logger.WithError(err).Fatal("Failed to start server")
	}
}
