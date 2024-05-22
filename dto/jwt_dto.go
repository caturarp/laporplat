package dto

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID     uint `json:"user_id"`
	PersonalID uint `json:"personal_id"`
	RoleID     uint `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(claims JWTClaims) (string, error) {
	expirationTime := time.Now().Add(15 * time.Hour)
	claims.Issuer = os.Getenv("APP_NAME")
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func GenerateRefreshToken(claims JWTClaims) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 30)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	claims.Issuer = os.Getenv("APP_NAME")
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
