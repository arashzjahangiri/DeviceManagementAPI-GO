package data

import (
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/config"
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database encapsulates database operations
type Database struct {
	DB  *gorm.DB
	log *logrus.Logger
}

// NewDatabase initializes a new Database instance
func NewDatabase(cfg *config.Config, log *logrus.Logger) (*Database, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the Device model
	if err := db.AutoMigrate(&models.Device{}); err != nil {
		return nil, err
	}

	return &Database{
		DB:  db,
		log: log,
	}, nil
}
