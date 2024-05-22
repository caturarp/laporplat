package dto

import (
	"time"
)

type UserParameter struct {
	ID    uint
	Email string `binding:"required,email" json:"email" validate:"required,email"`
}

type LoginRequest struct {
	Email    string `binding:"required,email" json:"email" validate:"required,email"`
	Password string `binding:"required" json:"password" validate:"required,min=8,max=35"`
}
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type RegisterRequest struct {
	Email      string `binding:"required,email" json:"email" validate:"required,email"`
	Password   string `binding:"required,min=8,max=30" json:"password" validate:"required,min=8,max=35"`
	VerifiedAt time.Time
	Code       string
}
type UpdateUserRequest struct {
	ID   uint    `json:"id"`
	Name *string `binding:"min=3,max=30" json:"name"`
}

type UserDetailResponse struct {
	Name  *string `json:"name"`
	Email string  `json:"email"`
}
type UserResponse struct {
	ID    uint    `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}
