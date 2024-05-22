package entity

type Vehicle struct {
	ID           uint   `json:"id"`
	LicensePlate string `json:"license_plate"` // non unique
	Year         int    `json:"year"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Color        string `json:"color"`
	VehicleType  string `json:"vehicle_type"`
}
