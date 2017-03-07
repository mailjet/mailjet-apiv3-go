package mailjet_test

import (
	"fmt"
	"net/textproto"
	"os"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/mailjet/mailjet-apiv3-go/resources"
)

var (
	publicKey  = os.Getenv("MJ_APIKEY_PUBLIC")
	privateKey = os.Getenv("MJ_APIKEY_PRIVATE")
)

func exampleMailjetClientList() {
	mj := mailjet.NewMailjetClient(publicKey, privateKey)

	var res []resources.Metadata
	count, total, err := mj.List("metadata", &res)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Count: %d\nTotal: %d\n", count, total)

	fmt.Println("Resources:")
	for _, resource := range res {
		fmt.Println(resource.Name)
	}
}

func exampleMailjetClientGet() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	var senders []resources.Sender
	info := &mailjet.Request{
		Resource: "sender",
		AltID:    "qwe@qwe.com",
	}
	err := mj.Get(info, &senders)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if senders != nil {
		fmt.Printf("Sender struct: %+v\n", senders[0])
	}
}

func exampleMailjetClientPost() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	var senders []resources.Sender
	fmr := &mailjet.FullRequest{
		Info:    &mailjet.Request{Resource: "sender"},
		Payload: &resources.Sender{Name: "Default", Email: "qwe@qwe.com"},
	}
	err := mj.Post(fmr, &senders)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if senders != nil {
		fmt.Printf("Data struct: %+v\n", senders[0])
	}
}

func exampleMailjetClientPut() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	fmr := &mailjet.FullRequest{
		Info:    &mailjet.Request{Resource: "sender", AltID: "qwe@qwe.com"},
		Payload: &resources.Sender{Name: "Bob", IsDefaultSender: true},
	}
	err := mj.Put(fmr, []string{"Name", "IsDefaultSender"})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println("Success")
	}
}

func exampleMailjetClientDelete() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	info := &mailjet.Request{
		Resource: "sender",
		AltID:    "qwe@qwe.com",
	}
	err := mj.Delete(info)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println("Success")
	}
}

func exampleMailjetClientSendMail() {
	publicKey := os.Getenv("MJ_APIKEY_PUBLIC")
	secretKey := os.Getenv("MJ_APIKEY_PRIVATE")

	mj := mailjet.NewMailjetClient(publicKey, secretKey)

	param := &mailjet.InfoSendMail{
		FromEmail: "qwe@qwe.com",
		FromName:  "Bob Patrick",
		Recipients: []mailjet.Recipient{
			{
				Email: "qwe@qwe.com",
			},
		},
		Subject:  "Hello World!",
		TextPart: "Hi there !",
	}
	res, err := mj.SendMail(param)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
		fmt.Println(res)
	}
}

func exampleMailjetClientSendMailSMTP() {
	mj := mailjet.NewMailjetClient(publicKey, privateKey)

	header := make(textproto.MIMEHeader)
	header.Add("From", "qwe@qwe.com")
	header.Add("To", "qwe@qwe.com")
	header.Add("Subject", "Hello World!")
	header.Add("X-Mailjet-Campaign", "test")
	content := []byte("Hi there !")
	info := &mailjet.InfoSMTP{
		From:       "qwe@qwe.com",
		Recipients: header["To"],
		Header:     header,
		Content:    content,
	}
	err := mj.SendMailSMTP(info)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
	}
}
