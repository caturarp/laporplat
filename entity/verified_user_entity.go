package entity

import "time"

type UnverifiedUser struct {
	ID        uint       `json:"id"`
	Email     string     `json:"email"`
	Code      string     `json:"code"`
	ExpiredAt time.Time  `json:"expired_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
