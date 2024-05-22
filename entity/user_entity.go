package entity

import "time"

type User struct {
	ID         uint       `json:"id"`
	Name       *string    `json:"name"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	VerifiedAt time.Time  `json:"verified_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

func (u User) IsEmpty() bool {
	return u == User{}
}
