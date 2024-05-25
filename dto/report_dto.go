package dto

type CreateReportRequest struct {
	Type            string `json:"type" binding:"required"`
	Description     string `json:"desc" binding:"required"`
	AreaOfOperation string `json:"area_of_operation" binding:"required"`
	Image           string `json:"image" binding:"required"`
	UserID          *uint  `json:"user_id"`
	DriverID        *uint  `json:"driver_id"`
	VehicleID       *uint  `json:"vehicle_id"`
	LicensePlate    string `json:"license_plate"`
}

type UpdateReportRequest struct {
	ID              uint    `json:"id" binding:"required"`
	Type            *string `json:"type" binding:"required"`
	Desc            *string `json:"desc" binding:"required"`
	AreaOfOperation *string `json:"area_of_operation" binding:"required"`
}
