package mailjet_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/mailjet/mailjet-apiv3-go/v4/resources"
)

var _ mailjet.ClientInterfaceV31 = (*mailjet.Client)(nil)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// defaultMessages is the default message passed to the server when an email
	// is sent.
	defaultMessages = mailjet.MessagesV31{
		Info: []mailjet.InfoMessagesV31{
			{
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
				TextPart: "SendMail is working!",
			},
		},
	}
)

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
		fmt.Fprint(w, response)
	})
}

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		//nolint:gosec // G404 crypto random is not required here
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
	tests := []struct {
		name           string
		messages       mailjet.MessagesV31
		mockResponse   interface{}
		mockStatusCode int
		wantResponse   *mailjet.ResultsV31
		wantErr        interface{}
	}{
		{
			name:     "sending successful",
			messages: defaultMessages,
			mockResponse: mailjet.ResultsV31{
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
			},
			mockStatusCode: 200,
			wantResponse: &mailjet.ResultsV31{
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
			},
			wantErr: nil,
		},
		{
			name:     "authorization failed",
			messages: defaultMessages,
			mockResponse: mailjet.ErrorInfoV31{
				Identifier: "ac93d194-1432-4e25-a215-2cb450d4a818",
				StatusCode: 401,
				Message:    "API key authentication/authorization failure. You may be unauthorized to access the API or your API key may be expired. Visit API keys management section to check your keys.",
			},
			mockStatusCode: 401,
			wantResponse:   nil,
			wantErr: &mailjet.ErrorInfoV31{
				Identifier: "ac93d194-1432-4e25-a215-2cb450d4a818",
				StatusCode: 401,
				Message:    "API key authentication/authorization failure. You may be unauthorized to access the API or your API key may be expired. Visit API keys management section to check your keys.",
			},
		},
		{
			name:     "simple errors in request",
			messages: messagesWithAdvancedErrorHandling(),
			mockResponse: mailjet.APIFeedbackErrorsV31{
				Messages: []mailjet.APIFeedbackErrorV31{
					{
						Errors: []mailjet.APIErrorDetailsV31{
							{
								ErrorCode:       "mj-0013",
								ErrorIdentifier: "ac93d194-1432-4e25-a215-2cb450d4a818",
								StatusCode:      400,
								// It is not, but let's suppose it is.
								ErrorMessage:   "\"recipient@company.com\" is an invalid email address.",
								ErrorRelatedTo: []string{"To[0].Email"},
							},
						},
					},
				},
			},
			mockStatusCode: 400,
			wantResponse:   nil,
			wantErr: &mailjet.APIFeedbackErrorsV31{
				Messages: []mailjet.APIFeedbackErrorV31{
					{
						Errors: []mailjet.APIErrorDetailsV31{
							{
								ErrorCode:       "mj-0013",
								ErrorIdentifier: "ac93d194-1432-4e25-a215-2cb450d4a818",
								StatusCode:      400,
								// It is not, but let's suppose it is.
								ErrorMessage:   "\"recipient@company.com\" is an invalid email address.",
								ErrorRelatedTo: []string{"To[0].Email"},
							},
						},
					},
				},
			},
		},
		{
			name:     "advanced error handling failed",
			messages: messagesWithAdvancedErrorHandling(),
			mockResponse: mailjet.APIFeedbackErrorsV31{
				Messages: []mailjet.APIFeedbackErrorV31{
					{
						Errors: []mailjet.APIErrorDetailsV31{
							{
								ErrorCode:       "send-0008",
								ErrorIdentifier: "ac93d194-1432-4e25-a215-2cb450d4a818",
								StatusCode:      403,
								ErrorMessage:    "\"passenger@mailjet.com\" is not an authorized sender email address for your account.",
								ErrorRelatedTo:  []string{"From"},
							},
						},
					},
				},
			},
			mockStatusCode: 403,
			wantResponse:   nil,
			wantErr: &mailjet.APIFeedbackErrorsV31{
				Messages: []mailjet.APIFeedbackErrorV31{
					{
						Errors: []mailjet.APIErrorDetailsV31{
							{
								ErrorCode:       "send-0008",
								ErrorIdentifier: "ac93d194-1432-4e25-a215-2cb450d4a818",
								StatusCode:      403,
								ErrorMessage:    "\"passenger@mailjet.com\" is not an authorized sender email address for your account.",
								ErrorRelatedTo:  []string{"From"},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// TODO(Go1.22+): remove:
			messages := test.messages // https://go.dev/wiki/CommonMistakes

			httpClientMocked := mailjet.NewhttpClientMock(true)
			httpClientMocked.SendMailV31Func = func(req *http.Request) (*http.Response, error) {
				if req.Header.Get("Content-Type") != "application/json" {
					t.Errorf("Wanted request content-type header to be: application/json, got: %s", req.Header.Get("Content-Type"))
				}

				user, pass, ok := req.BasicAuth()
				if !ok || user != httpClientMocked.APIKeyPublic() || pass != httpClientMocked.APIKeyPrivate() {
					t.Errorf("Wanted HTTP basic auth to be: %s/%s, got %s/%s", user, pass,
						httpClientMocked.APIKeyPublic(), httpClientMocked.APIKeyPrivate())
				}

				var msgs mailjet.MessagesV31
				err := json.NewDecoder(req.Body).Decode(&msgs)
				if err != nil {
					t.Fatalf("Could not decode request body and read message information: %v", err)
				}

				if !reflect.DeepEqual(test.messages, msgs) {
					t.Errorf("Wanted request messages: %+v, got: %+v", test.messages, msgs)
				}

				rawBytes, _ := json.Marshal(test.mockResponse)
				return &http.Response{
					Body:       ioutil.NopCloser(bytes.NewBuffer(rawBytes)),
					StatusCode: test.mockStatusCode,
				}, nil
			}

			m := mailjet.NewClient(httpClientMocked, mailjet.NewSMTPClientMock(true))

			res, err := m.SendMailV31(&messages)
			if !reflect.DeepEqual(err, test.wantErr) {
				t.Fatalf("Wanted error: %+v, got: %+v", err, test.wantErr)
			}

			if !reflect.DeepEqual(test.wantResponse, res) {
				t.Fatalf("Wanted response: %+v, got %+v", test.wantResponse, res)
			}
		})
	}
}

func messagesWithAdvancedErrorHandling() mailjet.MessagesV31 {
	withAdvancedErrorChecking := defaultMessages
	withAdvancedErrorChecking.AdvanceErrorHandling = true
	return withAdvancedErrorChecking
}
