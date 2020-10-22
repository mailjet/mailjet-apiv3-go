package mailjet_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mailjet "github.com/mailjet/mailjet-apiv3-go/v3"
	"github.com/mailjet/mailjet-apiv3-go/v3/resources"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *mailjet.Client
)

func fakeServer() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = mailjet.NewMailjetClient("apiKeyPublic", "apiKeyPrivate", server.URL+"/v3")

	return func() {
		server.Close()
	}
}

func handle(path, response string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, response)
	})
}

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

func TestCreateListrecipient(t *testing.T) {
	teardown := fakeServer()
	defer teardown()

	mux.HandleFunc("/v3/REST/listrecipient", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal("Unexpected error:", err)
		}

		body := make(map[string]interface{})
		if err = json.Unmarshal(b, &body); err != nil {
			t.Fatal("Invalid body:", err)
		}

		_, id := body["ContactID"]
		_, alt := body["ContactALT"]
		_, listID := body["ListID"]
		if !id && !alt || !listID {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `{"ErrorMessage": "Missing required parameters"}`)
		}
	})

	t.Run("successfully create list", func(t *testing.T) {
		req := &mailjet.Request{
			Resource: "listrecipient",
		}
		fullRequest := &mailjet.FullRequest{
			Info: req,
			Payload: resources.Listrecipient{
				IsUnsubscribed: true,
				ContactID:      124409882,
				ContactALT:     "joe.doe@mailjet.com",
				ListID:         32964,
			},
		}

		var resp []resources.Listrecipient
		err := client.Post(fullRequest, &resp)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("failure when required parameters missing", func(t *testing.T) {
		req := &mailjet.Request{
			Resource: "listrecipient",
		}
		fullRequest := &mailjet.FullRequest{
			Info: req,
			Payload: resources.Listrecipient{
				IsUnsubscribed: true,
				ListID:         32964,
			},
		}

		var resp []resources.Listrecipient
		err := client.Post(fullRequest, &resp)
		if err == nil {
			t.Fatal("Expected error")
		}
	})
}

func TestMessage(t *testing.T) {
	teardown := fakeServer()
	defer teardown()

	handle("/v3/REST/message", `
	{
		"Count": 1,
		"Data": [
			{
				"ArrivedAt": "2020-10-08T06:36:35Z",
				"AttachmentCount": 0,
				"AttemptCount": 0,
				"CampaignID": 426400,
				"ContactAlt": "",
				"ContactID": 124409882,
				"Delay": 0,
				"DestinationID": 124879,
				"FilterTime": 0,
				"ID": 94294117474376580,
				"IsClickTracked": false,
				"IsHTMLPartIncluded": false,
				"IsOpenTracked": true,
				"IsTextPartIncluded": false,
				"IsUnsubTracked": false,
				"MessageSize": 810,
				"SenderID": 52387,
				"SpamassassinScore": 0,
				"SpamassRules": "",
				"StatePermanent": false,
				"Status": "sent",
				"Subject": "",
				"UUID": "6f66806a-c4d6-4a33-99dc-bedbc7c4217f"
			}
		],
		"Total": 1
	}
   `)

	request := &mailjet.Request{
		Resource: "message",
	}

	var data []resources.Message

	err := client.Get(request, &data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMessageinformation(t *testing.T) {
	t.Run("empty SpamAssassinRules", func(t *testing.T) {
		teardown := fakeServer()
		defer teardown()

		handle("/v3/REST/messageinformation", `
		{
			"Count": 1,
			"Data": [
				{
					"CampaignID": 0,
					"ClickTrackedCount": 0,
					"ContactID": 124409882,
					"CreatedAt": "2020-10-09T06:07:56Z",
					"ID": 288230380871887400,
					"MessageSize": 434,
					"OpenTrackedCount": 0,
					"QueuedCount": 0,
					"SendEndAt": "2020-10-09T06:07:56Z",
					"SentCount": 1602223677,
					"SpamAssassinRules": {
						"ALT": "",
						"ID": -1
					},
					"SpamAssassinScore": 0
				}
			],
			"Total": 1
		}
		`)

		request := &mailjet.Request{
			Resource: "messageinformation",
		}

		var data []resources.Messageinformation

		err := client.Get(request, &data)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("not empty SpamAssassinRules", func(t *testing.T) {
		teardown := fakeServer()
		defer teardown()

		handle("/v3/REST/messageinformation", `
		{
			"Count": 1,
			"Data": [
				{
					"CampaignID": 0,
					"ClickTrackedCount": 0,
					"ContactID": 124409882,
					"CreatedAt": "2020-10-09T06:07:56Z",
					"ID": 288230380871887400,
					"MessageSize": 434,
					"OpenTrackedCount": 0,
					"QueuedCount": 0,
					"SendEndAt": "2020-10-09T06:07:56Z",
					"SentCount": 1602223677,
					"SpamAssassinRules": {
						"ALT": "",
						"ID": -1,
						"Items": [
							{
								"ALT": "MISSING_DATE",
								"HitCount": 81115,
								"ID": 1,
								"Name": "MISSING_DATE",
								"Score": 2.739
							},
							{
								"ALT": "MISSING_HEADERS",
								"HitCount": 48433743,
								"ID": 2,
								"Name": "MISSING_HEADERS",
								"Score": 0.915
							}
						]
					},
					"SpamAssassinScore": 0
				}
			],
			"Total": 1
		}
		`)

		request := &mailjet.Request{
			Resource: "messageinformation",
		}

		var data []resources.Messageinformation

		err := client.Get(request, &data)
		if err != nil {
			t.Fatal(err)
		}
	})
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
