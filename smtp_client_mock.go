package mailjet

import (
	"net/smtp"
)

// smtpClientMock def
type smtpClientMock struct {
}

// SendMail wraps smtp.SendMail
func (s smtpClientMock) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return nil
}
