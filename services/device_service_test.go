package services

import (
	"testing"

	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/data"
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/dtos"
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestDeviceService_GetAllDevices tests the GetAllDevices method
func TestDeviceService_GetAllDevices(t *testing.T) {
	// Setup in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.Device{})
	assert.NoError(t, err)

	// Create test data
	devices := []models.Device{
		{Name: "Device1", Type: "Sensor"},
		{Name: "Device2", Type: "Actuator"},
	}
	for _, device := range devices {
		err := db.Create(&device).Error
		assert.NoError(t, err)
	}

	// Initialize service
	database := &data.Database{DB: db, log: logrus.New()}
	service := NewDeviceService(database, logrus.New())

	// Test
	result, err := service.GetAllDevices()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Device1", result[0].Name)
	assert.Equal(t, "Sensor", result[0].Type)
	assert.Equal(t, "Device2", result[1].Name)
	assert.Equal(t, "Actuator", result[1].Type)
}

// TestDeviceService_CreateDevice tests the CreateDevice method
func TestDeviceService_CreateDevice(t *testing.T) {
	// Setup in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.Device{})
	assert.NoError(t, err)

	// Initialize service
	database := &data.Database{DB: db, log: logrus.New()}
	service := NewDeviceService(database, logrus.New())

	// Test
	req := dtos.CreateDeviceRequest{Name: "Device1", Type: "Sensor"}
	deviceDto, err := service.CreateDevice(req)
	assert.NoError(t, err)
	assert.NotNil(t, deviceDto)
	assert.Equal(t, "Device1", deviceDto.Name)
	assert.Equal(t, "Sensor", deviceDto.Type)

	// Verify in database
	var device models.Device
	err = db.First(&device, deviceDto.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, "Device1", device.Name)
	assert.Equal(t, "Sensor", device.Type)
}

// TestDeviceService_UpdateDevice_NotFound tests the UpdateDevice method when the device doesn't exist
func TestDeviceService_UpdateDevice_NotFound(t *testing.T) {
	// Setup in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&models.Device{})
	assert.NoError(t, err)

	// Initialize service
	database := &data.Database{DB: db, log: logrus.New()}
	service := NewDeviceService(database, logrus.New())

	// Test
	req := dtos.UpdateDeviceRequest{Name: "UpdatedDevice", Type: "Actuator"}
	success, err := service.UpdateDevice(999, req)
	assert.Error(t, err)
	assert.False(t, success)
}
