package mailjet

import (
	"bytes"
	"fmt"
	"net/smtp"
	"net/textproto"
	"strings"
)

// InfoSMTP contains mandatory informations to send a mail via SMTP.
type InfoSMTP struct {
	From       string
	Recipients []string
	Header     textproto.MIMEHeader
	Content    []byte
}

// Hostname and port for the SMTP client.
const (
	HostSMTP = "in-v3.mailjet.com"
	PortSMTP = 587
)

// SendMailSMTP send mail via SMTP.
func (mj *Client) SendMailSMTP(info *InfoSMTP) (err error) {
	auth := smtp.PlainAuth(
		"",
		mj.APIKeyPublic(),
		mj.APIKeyPrivate(),
		HostSMTP,
	)

	host := fmt.Sprintf("%s:%d", HostSMTP, PortSMTP)
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
