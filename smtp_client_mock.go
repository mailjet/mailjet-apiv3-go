package mailjet

import (
	"errors"
)

// SMTPClientMock def
type SMTPClientMock struct {
	valid bool
}

// NewSMTPClientMock returns a new smtp client mock
func NewSMTPClientMock(valid bool) *SMTPClientMock {
	return &SMTPClientMock{
		valid: valid,
	}
}

// SendMail wraps smtp.SendMail
func (s SMTPClientMock) SendMail(from string, to []string, msg []byte) error {
	if s.valid == true {
		return nil
	}
	return errors.New("smtp send error")
}
