package mailjet

import (
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go/resources"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"
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
	m := NewMailjetClient(ak, sk)

	if ak != m.ApiKeyPublic() {
		t.Fatal("Wrong public key:", m.ApiKeyPublic())
	}

	if sk != m.ApiKeyPrivate() {
		t.Fatal("Wrong secret key:", m.ApiKeyPrivate())
	}

	if http.DefaultClient != m.Client() {
		t.Fatal("HTTP client not default!")
	}
	client := new(http.Client)
	m.SetClient(client)
	if client != m.Client() {
		t.Fatal("HTTP client not equal!")
	}
}

func TestList(t *testing.T) {
	m := NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

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
	m := NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

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

	mr := &MailjetRequest{Resource: resource, ID: data[0].ID}
	data = make([]resources.User, 0)
	err = m.Get(mr, &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	fmt.Printf("Data: %+v\n", data[0])
}

func TestPost(t *testing.T) {
	m := NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

	var data []resources.Contact
	rstr := randSeq(10)
	fmt.Printf("Create new contact: \"%s@mailjet.com\"\n", rstr)
	fmr := &FullMailjetRequest{
		Info:    &MailjetRequest{Resource: "contact"},
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
	m := NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

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
	fmr := &FullMailjetRequest{
		Info:    &MailjetRequest{Resource: resource, AltID: data[0].Address},
		Payload: data[0],
	}
	err = m.Put(fmr, []string{"Name"})
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

}

func TestDelete(t *testing.T) {
	m := NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

	var data []resources.Listrecipient
	resource := "listrecipient"
	count, _, err := m.List(resource, &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if count < 1 {
		t.Fatal("At least one listrecipient expected !")
	}
	if data == nil {
		t.Fatal("Empty result")
	}

	mr := &MailjetRequest{
		ID:       data[0].ID,
		Resource: resource,
	}
	err = m.Delete(mr)
	if err != nil {
		fmt.Println(err)
	}
}

func TestSendMail(t *testing.T) {
	m := NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))

	var data []resources.Sender
	count, _, err := m.List("sender", &data)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}
	if count < 1 || data == nil {
		t.Fatal("At least one sender expected in the test account!")
	}

	param := &MailjetSendMail{
		FromEmail: data[0].Email,
		FromName:  data[0].Name,
		Recipients: []MailjetRecipient{
			MailjetRecipient{
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
