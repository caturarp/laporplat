package mailer

import (
	"fmt"
	"net/smtp"
	"os"
)

type SMTPConfig struct {
	Host       string
	Port       string
	SenderName string
	Email      string
	Password   string
}

func (s *SMTPConfig) AddressBuilder() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

func GetSMTPConfig() SMTPConfig {
	smtpHost := os.Getenv("CONFIG_SMTP_HOST")
	smtpPort := os.Getenv("CONFIG_SMTP_PORT")
	senderName := os.Getenv("CONFIG_SENDER_NAME")
	authEmail := os.Getenv("CONFIG_AUTH_EMAIL")
	authPassword := os.Getenv("CONFIG_AUTH_PASSWORD")
	return SMTPConfig{
		Host:       smtpHost,
		Port:       smtpPort,
		SenderName: senderName,
		Email:      authEmail,
		Password:   authPassword,
	}
}
func InitAuthSMTP() smtp.Auth {
	smtpConfig := GetSMTPConfig()
	auth := smtp.PlainAuth("", smtpConfig.Email, smtpConfig.Password, smtpConfig.Host)
	return auth
}
