package dto

type SendVerificationMailRequest struct {
	EmailRecipients []string `binding:"required,email" json:"email_recipients"`
	EmailCCs        []string `binding:"required,email" json:"email_ccs"`
	Subject         string   `json:"subject"`
	Content         string   `json:"content"`
}
