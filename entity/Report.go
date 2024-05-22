package entity

import "time"

type Report struct {
	ID           uint       `json:"id"`
	VehicleID    uint       `json:"vehicle_id"`
	UserID       uint       `json:"user_id"`
	LicensePlate string     `json:"license_plate"`
	ReportType   string     `json:"report_type"`
	Description  string     `json:"description"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
