package mailjet

import "net/smtp"

// smtpClient def
type smtpClient struct {
}

// SendMail wraps smtp.SendMail
func (s smtpClient) SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	return smtp.SendMail(addr, a, from, to, msg)
}
