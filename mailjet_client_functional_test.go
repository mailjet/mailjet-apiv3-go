package mailjet_test

import (
	"fmt"
	"math/rand"
	"net/textproto"
	"os"
	"sync"
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

func TestNewMailjetClient(t *testing.T) {
	ak := os.Getenv("MJ_APIKEY_PUBLIC")
	sk := os.Getenv("MJ_APIKEY_PRIVATE")
	m := mailjet.NewMailjetClient(ak, sk)

	if ak != m.APIKeyPublic() {
		t.Fatal("Wrong public key:", m.APIKeyPublic())
	}

	if sk != m.APIKeyPrivate() {
		t.Fatal("Wrong secret key:", m.APIKeyPrivate())
	}
}

func TestList(t *testing.T) {
	m := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

	var data []resources.Sender
	count, _, err := m.List("sender", &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if count < 1 {
		t.Fatal("At least one sender expected !")
	}
}

func TestGet(t *testing.T) {
	m := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

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
	fmt.Printf("Data: %+v\n", data[0])
}

func TestPost(t *testing.T) {
	m := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

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

func TestPut(t *testing.T) {
	m := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

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

func TestDelete(t *testing.T) {
	m := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

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

func TestSendMail(t *testing.T) {
	m := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

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
			mailjet.Recipient{
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

func TestSendMailSmtp(t *testing.T) {
	mj := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

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
	info := &mailjet.InfoSMTP{
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

func TestDataRace(t *testing.T) {
	m := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			var data []resources.Sender
			count, _, err := m.List("sender", &data)
			if err != nil {
				t.Fatal("Unexpected error:", err)
			}
			if count < 1 || data == nil {
				t.Fatal("At least one sender expected in the test account!")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
