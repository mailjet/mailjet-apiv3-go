package mailjet_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
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
	// Here we set the behavior of the http mock
	httpClientMocked := mailjet.NewhttpClientMock(true)
	httpClientMocked.SendMailV31Func = func(req *http.Request) (*http.Response, error) {
		data := mailjet.ResultsV31{
			ResultsV31: []mailjet.ResultV31{
				{
					To: []mailjet.GeneratedMessageV31{
						{
							Email:       "recipient@company.com",
							MessageUUID: "ac93d194-1432-4e25-a215-2cb450d4a818",
							MessageID:   87,
						},
					},
				},
			},
		}
		rawBytes, _ := json.Marshal(data)
		return &http.Response{
			Body:       ioutil.NopCloser(bytes.NewBuffer(rawBytes)),
			StatusCode: http.StatusOK,
		}, nil
	}

	m := mailjet.NewClient(httpClientMocked, mailjet.NewSMTPClientMock(true))

	// We define parameters here to pass to SendMailV31
	param := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "passenger@mailjet.com",
				Name:  "passenger",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: "recipient@company.com",
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
