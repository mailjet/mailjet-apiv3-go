package mailjet_test

import (
	"fmt"
	"testing"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailjet/mailjet-apiv3-go/resources"
)

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
	fmt.Printf("Create new contact: \"%s@mailjet.com\"\n", rstr)
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
	fmt.Printf("Data: %+v\n", data[0])
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
	fmt.Printf("Update name of the contact list: %s -> %s\n", data[0].Name, rstr)
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
		fmt.Println(err)
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
	res, err := m.SendMail(param)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	fmt.Printf("Data: %+v\n", res)
}
