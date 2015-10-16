package mailjet

import (
	"bytes"
	"fmt"
	"net/smtp"
	"net/textproto"
	"strings"
)

// MailjetSmtp contains mandatory informations to send a mail via SMTP.
type MailjetSmtp struct {
	From       string
	Recipients []string
	Header     textproto.MIMEHeader
	Content    []byte
}

const (
	MailjetHostSmtp = "in-v3.mailjet.com"
	MailjetPortSmtp = 587
)

// SendMailSmtp send mail via SMTP.
func (mj *MailjetClient) SendMailSmtp(info *MailjetSmtp) (err error) {
	auth := smtp.PlainAuth(
		"",
		mj.apiKeyPublic,
		mj.apiKeyPrivate,
		MailjetHostSmtp,
	)

	host := fmt.Sprintf("%s:%d", MailjetHostSmtp, MailjetPortSmtp)
	err = smtp.SendMail(
		host,
		auth,
		info.From,
		info.Recipients,
		buildMessage(info.Header, info.Content),
	)

	return err
}

func buildMessage(header textproto.MIMEHeader, content []byte) []byte {
	buff := bytes.NewBuffer(nil)
	for key, values := range header {
		buff.WriteString(fmt.Sprintf("%s: %s\r\n", key, strings.Join(values, ", ")))
	}
	buff.WriteString("\r\n")
	buff.Write(content)

	return buff.Bytes()
}
