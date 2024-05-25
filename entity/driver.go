package entity

type Driver struct {
	ID            int     `json:"id"`             // Unique identifier (primary key in the database)
	Name          string  `json:"name"`           // Full name of the driver
	LicenseNumber string  `json:"license_number"` // Driver's license number
	RideService   string  `json:"ride_service"`   // Ride-hailing service (e.g., "Uber", "Lyft")
	PhoneNumber   string  `json:"phone_number"`   // Driver's phone number (if applicable)
	Email         string  `json:"email"`          // Driver's email address (if applicable)
	Rating        float32 `json:"rating"`         // Average rating from users (if applicable)
	NumReports    int     `json:"-"`              // Number of reports filed against the driver (not exposed in JSON)
	IsActive      bool    `json:"-"`              // Whether the driver is currently active
	// ... other relevant fields (e.g., photo, license expiration date, etc.)
}
