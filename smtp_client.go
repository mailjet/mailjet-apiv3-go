package mailjet

import (
	"fmt"
	"net/smtp"
)

// SMTPClient is the wrapper for smtp
type SMTPClient struct {
	host string
	auth smtp.Auth
}

// Hostname and port for the SMTP client.
const (
	HostSMTP = "in-v3.mailjet.com"
	PortSMTP = 587
)

// NewSMTPClient returns a new smtp client wrapper
func NewSMTPClient(apiKeyPublic, apiKeyPrivate string) *SMTPClient {
	auth := smtp.PlainAuth(
		"",
		apiKeyPublic,
		apiKeyPrivate,
		HostSMTP,
	)
	return &SMTPClient{
		host: fmt.Sprintf("%s:%d", HostSMTP, PortSMTP),
		auth: auth,
	}
}

// SendMail wraps smtp.SendMail
func (s SMTPClient) SendMail(from string, to []string, msg []byte) error {
	return smtp.SendMail(s.host, s.auth, from, to, msg)
}
