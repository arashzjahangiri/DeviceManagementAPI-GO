package controllers

import (
	"net/http"
	"strconv"

	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/dtos"
	"github.com/arashzjahangiri/DeviceManagementAPI-GO/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// DevicesController handles HTTP requests for devices
type DevicesController struct {
	service *services.DeviceService
	log     *logrus.Logger
}

// NewDevicesController creates a new DevicesController instance
func NewDevicesController(service *services.DeviceService, log *logrus.Logger) *DevicesController {
	return &DevicesController{
		service: service,
		log:     log,
	}
}

// GetDevices handles GET /devices
func (c *DevicesController) GetDevices(ctx *gin.Context) {
	c.log.Debug("Handling GET /devices request")
	devices, err := c.service.GetAllDevices()
	if err != nil {
		c.log.WithError(err).Error("Failed to fetch devices")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, devices)
}

// GetDevice handles GET /devices/:id
func (c *DevicesController) GetDevice(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.log.WithField("id", idStr).Warn("Invalid device ID")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	c.log.WithField("id", id).Debug("Handling GET /devices/:id request")
	device, err := c.service.GetDeviceByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.log.WithField("id", id).Warn("Device not found")
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
			return
		}
		c.log.WithError(err).WithField("id", id).Error("Failed to fetch device")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, device)
}

// CreateDevice handles POST /devices
func (c *DevicesController) CreateDevice(ctx *gin.Context) {
	c.log.Debug("Handling POST /devices request")
	var req dtos.CreateDeviceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.log.WithError(err).Warn("Invalid request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	device, err := c.service.CreateDevice(req)
	if err != nil {
		c.log.WithError(err).Error("Failed to create device")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusCreated, device)
}

// UpdateDevice handles PUT /devices/:id
func (c *DevicesController) UpdateDevice(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.log.WithField("id", idStr).Warn("Invalid device ID")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	c.log.WithField("id", id).Debug("Handling PUT /devices/:id request")
	var req dtos.UpdateDeviceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.log.WithError(err).Warn("Invalid request body")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := c.service.UpdateDevice(uint(id), req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.log.WithField("id", id).Warn("Device not found")
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
			return
		}
		c.log.WithError(err).WithField("id", id).Error("Failed to update device")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if !success {
		c.log.WithField("id", id).Error("Failed to update device")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// DeleteDevice handles DELETE /devices/:id
func (c *DevicesController) DeleteDevice(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.log.WithField("id", idStr).Warn("Invalid device ID")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	c.log.WithField("id", id).Debug("Handling DELETE /devices/:id request")
	success, err := c.service.DeleteDevice(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.log.WithField("id", id).Warn("Device not found")
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
			return
		}
		c.log.WithError(err).WithField("id", id).Error("Failed to delete device")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if !success {
		c.log.WithField("id", id).Error("Failed to delete device")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
