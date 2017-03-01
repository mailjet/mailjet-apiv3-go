package mailjet

import "net/smtp"

// smtpClientInterface def
type smtpClientInterface interface {
	SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error
}
