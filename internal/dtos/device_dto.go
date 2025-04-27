package dtos

// CreateDeviceRequest represents the request body for creating a device
type CreateDeviceRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

// UpdateDeviceRequest represents the request body for updating a device
type UpdateDeviceRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

// DeviceDto represents the response DTO for a device
type DeviceDto struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
