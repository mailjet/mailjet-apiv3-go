package mailjet

import (
	"net/textproto"
	"os"
	"testing"

	"github.com/mailjet/mailjet-apiv3-go/resources"
)

func TestSendMailSmtp(t *testing.T) {
	mj := NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

	var data []resources.Sender
	count, _, err := mj.List("sender", &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if count < 1 || data == nil {
		t.Fatal("At least one sender expected in the test account!")
	}

	email := data[0].Email

	header := make(textproto.MIMEHeader)
	header.Add("From", email)
	header.Add("To", email)
	header.Add("Subject", "SMTP testing")
	header.Add("X-Mailjet-Campaign", "test")
	content := []byte("SendMailSmtp is working !")
	info := &InfoSMTP{
		From:       email,
		Recipients: header["To"],
		Header:     header,
		Content:    content,
	}
	err = mj.SendMailSMTP(info)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
}
