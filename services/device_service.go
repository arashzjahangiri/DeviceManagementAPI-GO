package services

import (
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/data"
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/dtos"
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// DeviceService handles business logic for devices
type DeviceService struct {
	db  *data.Database
	log *logrus.Logger
}

// NewDeviceService creates a new DeviceService instance
func NewDeviceService(db *data.Database, log *logrus.Logger) *DeviceService {
	return &DeviceService{
		db:  db,
		log: log,
	}
}

// GetAllDevices retrieves all devices
func (s *DeviceService) GetAllDevices() ([]dtos.DeviceDto, error) {
	s.log.Debug("Fetching all devices")
	var devices []models.Device
	if err := s.db.DB.Find(&devices).Error; err != nil {
		s.log.WithError(err).Error("Failed to fetch devices")
		return nil, err
	}

	deviceDtos := make([]dtos.DeviceDto, len(devices))
	for i, device := range devices {
		deviceDtos[i] = dtos.DeviceDto{
			ID:   device.ID,
			Name: device.Name,
			Type: device.Type,
		}
	}
	s.log.Debugf("Found %d devices", len(deviceDtos))
	return deviceDtos, nil
}

// GetDeviceByID retrieves a device by ID
func (s *DeviceService) GetDeviceByID(id uint) (*dtos.DeviceDto, error) {
	s.log.WithField("id", id).Debug("Fetching device by ID")
	var device models.Device
	if err := s.db.DB.First(&device, id).Error; err != nil {
		s.log.WithError(err).WithField("id", id).Warn("Device not found")
		return nil, err
	}

	deviceDto := &dtos.DeviceDto{
		ID:   device.ID,
		Name: device.Name,
		Type: device.Type,
	}
	s.log.WithField("id", id).Debug("Device fetched successfully")
	return deviceDto, nil
}

// CreateDevice creates a new device
func (s *DeviceService) CreateDevice(req dtos.CreateDeviceRequest) (*dtos.DeviceDto, error) {
	s.log.WithFields(logrus.Fields{
		"name": req.Name,
		"type": req.Type,
	}).Debug("Creating new device")

	device := models.Device{
		Name: req.Name,
		Type: req.Type,
	}
	if err := s.db.DB.Create(&device).Error; err != nil {
		s.log.WithError(err).Error("Failed to create device")
		return nil, err
	}

	deviceDto := &dtos.DeviceDto{
		ID:   device.ID,
		Name: device.Name,
		Type: device.Type,
	}
	s.log.WithField("id", device.ID).Info("Device created successfully")
	return deviceDto, nil
}

// UpdateDevice updates an existing device
func (s *DeviceService) UpdateDevice(id uint, req dtos.UpdateDeviceRequest) (bool, error) {
	s.log.WithField("id", id).Debug("Updating device")
	var device models.Device
	if err := s.db.DB.First(&device, id).Error; err != nil {
		s.log.WithError(err).WithField("id", id).Warn("Device not found for update")
		return false, err
	}

	device.Name = req.Name
	device.Type = req.Type
	if err := s.db.DB.Save(&device).Error; err != nil {
		s.log.WithError(err).WithField("id", id).Error("Failed to update device")
		return false, err
	}

	s.log.WithField("id", id).Info("Device updated successfully")
	return true, nil
}

// DeleteDevice deletes a device by ID
func (s *DeviceService) DeleteDevice(id uint) (bool, error) {
	s.log.WithField("id", id).Debug("Deleting device")
	result := s.db.DB.Delete(&models.Device{}, id)
	if result.Error != nil {
		s.log.WithError(err).WithField("id", id).Error("Failed to delete device")
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		s.log.WithField("id", id).Warn("Device not found for deletion")
		return false, gorm.ErrRecordNotFound
	}

	s.log.WithField("id", id).Info("Device deleted successfully")
	return true, nil
}
