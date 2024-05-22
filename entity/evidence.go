package entity

import "time"

type ImageEvidence struct {
	ID        uint      `json:"id"`
	ReportID  uint      `json:"report_id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VideoEvidence struct {
	ID        uint      `json:"id"`
	ReportID  uint      `json:"report_id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
