package util

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/caturarp/laporplat/dto"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	password := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	return err == nil
}
func GenerateCode(registerInfo dto.VerifyRequest) string {
	rand.Seed(time.Now().UnixNano())

	randomString := func(length int) string {
		charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
		b := make([]byte, length)
		for i := range b {
			b[i] = charset[rand.Intn(len(charset))]
		}
		return string(b)
	}

	hasher := sha256.New()
	hasher.Write([]byte(randomString(10)))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash
}
