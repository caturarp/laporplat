package mail

import (
	"context"
	"net/smtp"
	"strings"

	"github.com/caturarp/laporplat/dto"
	mailer "github.com/caturarp/laporplat/external/config"
)

type SMTPMailer interface {
	SendMail(context.Context, dto.SendVerificationMailRequest) error
}

type smtpMailer struct {
	smtpConfig mailer.SMTPConfig
	auth       smtp.Auth
}

func NewSMTP(smtpConfig mailer.SMTPConfig) SMTPMailer {
	auth := mailer.InitAuthSMTP()
	return &smtpMailer{
		smtpConfig: smtpConfig,
		auth:       auth,
	}
}

func (s *smtpMailer) SendMail(ctx context.Context, mailDetail dto.SendVerificationMailRequest) error {
	body := "From: " + s.smtpConfig.SenderName + "\n" +
		"To: " + strings.Join(mailDetail.EmailRecipients, ",") + "\n" +
		"Cc: " + strings.Join(mailDetail.EmailCCs, ",") + "\n" +
		"Subject: " + mailDetail.Subject + "\n\n" +
		mailDetail.Content
	smtpAddress := s.smtpConfig.AddressBuilder()
	err := smtp.SendMail(smtpAddress, s.auth, s.smtpConfig.Email, append(mailDetail.EmailRecipients, mailDetail.EmailCCs...), []byte(body))
	if err != nil {
		return err
	}
	return nil
}
