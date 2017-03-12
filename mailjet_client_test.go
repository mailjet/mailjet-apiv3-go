package mailjet_test

import (
	"math/rand"
	"os"
	"testing"
	"time"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailjet/mailjet-apiv3-go/resources"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// NewMockedMailjetClient returns an instance of `Client` with mocked http and smtp clients injected
func newMockedMailjetClient() *mailjet.Client {
	httpClientMocked := mailjet.NewhttpClientMock(true)
	smtpClientMocked := mailjet.NewSMTPClientMock(true)
	client := mailjet.NewClient(httpClientMocked, smtpClientMocked)

	return client
}

func TestUnitList(t *testing.T) {
	m := newMockedMailjetClient()

	var data []resources.Sender
	count, _, err := m.List("sender", &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if count < 1 {
		t.Fatal("At least one sender expected !")
	}

	httpClientMocked := mailjet.NewhttpClientMock(false)
	smtpClientMocked := mailjet.NewSMTPClientMock(true)
	cl := mailjet.NewClient(httpClientMocked, smtpClientMocked, "custom")

	_, _, err = cl.List("sender", &data)
	if err == nil {
		t.Fail()
	}
}

func TestUnitGet(t *testing.T) {
	m := newMockedMailjetClient()

	var data []resources.User
	resource := "user"
	count, _, err := m.List(resource, &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if count < 1 {
		t.Fatal("At least one user expected !")
	}
	if data == nil {
		t.Fatal("Empty result")
	}

	mr := &mailjet.Request{Resource: resource, ID: data[0].ID}
	data = make([]resources.User, 0)
	err = m.Get(mr, &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
}

func TestUnitPost(t *testing.T) {
	m := newMockedMailjetClient()

	var data []resources.Contact
	rstr := randSeq(10)
	t.Logf("Create new contact: \"%s@mailjet.com\"\n", rstr)
	fmr := &mailjet.FullRequest{
		Info:    &mailjet.Request{Resource: "contact"},
		Payload: &resources.Contact{Name: rstr, Email: rstr + "@mailjet.com"},
	}
	err := m.Post(fmr, &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if data == nil {
		t.Fatal("Empty result")
	}
	t.Logf("Created contact: %+v\n", data[0])
}

func TestUnitPut(t *testing.T) {
	m := newMockedMailjetClient()

	var data []resources.Contactslist
	resource := "contactslist"
	count, _, err := m.List(resource, &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if count < 1 {
		t.Fatal("At least one contact list expected on test account!")
	}
	if data == nil {
		t.Fatal("Empty result")
	}

	rstr := randSeq(10)
	t.Logf("Update name of the contact list: %s -> %s\n", data[0].Name, rstr)
	data[0].Name = randSeq(10)
	fmr := &mailjet.FullRequest{
		Info:    &mailjet.Request{Resource: resource, AltID: data[0].Address},
		Payload: data[0],
	}
	err = m.Put(fmr, []string{"Name"})
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
}

func TestUnitDelete(t *testing.T) {
	m := newMockedMailjetClient()

	var data []resources.Listrecipient
	resource := "listrecipient"
	count, _, err := m.List(resource, &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if count < 1 {
		return
	}
	if data == nil {
		t.Fatal("Empty result")
	}

	mr := &mailjet.Request{
		ID:       data[0].ID,
		Resource: resource,
	}
	err = m.Delete(mr)
	if err != nil {
		t.Error(err)
	}
}

func TestUnitSendMail(t *testing.T) {
	m := newMockedMailjetClient()

	var data []resources.Sender
	count, _, err := m.List("sender", &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if count < 1 || data == nil {
		t.Fatal("At least one sender expected in the test account!")
	}

	param := &mailjet.InfoSendMail{
		FromEmail: data[0].Email,
		FromName:  data[0].Name,
		Recipients: []mailjet.Recipient{
			{
				Email: data[0].Email,
			},
		},
		Subject:  "Send API testing",
		TextPart: "SendMail is working !",
	}
	_, err = m.SendMail(param)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
}

func TestSendMailV31(t *testing.T) {
	m := mailjet.NewMailjetClient(
		os.Getenv("MJ_APIKEY_PUBLIC"),
		os.Getenv("MJ_APIKEY_PRIVATE"))

	var data []resources.Sender
	count, _, err := m.List("sender", &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if count < 1 || data == nil {
		t.Fatal("At least one sender expected in the test account!")
	}

	param := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: data[0].Email,
				Name:  data[0].Name,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: data[0].Email,
				},
			},
			Subject:  "Send API testing",
			TextPart: "SendMail is working !",
		},
	}

	messages := mailjet.MessagesV31{Info: param}

	res, err := m.SendMailV31(&messages)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Logf("Data: %+v\n", res)
	}
}
