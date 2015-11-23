package mailjet

import (
	"net/http"
)

// MailjetClient bundles data needed by a large number
// of methods in order to interact with the Mailjet API.
type MailjetClient struct {
	apiKeyPublic  string
	apiKeyPrivate string
	client        *http.Client
}

// MailjetRequest bundles data needed to build the URL.
type MailjetRequest struct {
	Resource string
	ID       int64
	AltID    string
	Action   string
	ActionID int64
}

// MailjetDataRequest bundles data needed to build the DATA URL.
type MailjetDataRequest struct {
	SourceType   string
	SourceTypeID int64
	DataType     string
	MimeType     string
	DataTypeID   int64
	LastID       bool
}

// FullMailjetRequest is the same as a MailjetRequest but with a payload.
type FullMailjetRequest struct {
	Info    *MailjetRequest
	Payload interface{}
}

// FullMailjetDataRequest is the same as a MailjetDataRequest but with a payload.
type FullMailjetDataRequest struct {
	Info    *MailjetDataRequest
	Payload interface{}
}

type MailjetOptions func(*http.Request)

// MailjetResult is the JSON result sent by the API.
type MailjetResult struct {
	Count int
	Data  interface{}
	Total int
}

// MailjetError is the error returned by the API.
type MailjetError struct {
	ErrorInfo    string
	ErrorMessage string
	StatusCode   int
}

//
// Send API structures
//

type MailjetSendMail struct {
	FromEmail             string
	FromName              string
	Sender                string
	Recipients            []MailjetRecipient
	To                    []string
	Cc                    []string
	Bcc                   []string
	Subject               string
	TextPart              string `json:"Text-part"`
	HtmlPart              string `json:"Html-part"`
	Attachments           []MailjetAttachment
	InlineAttachments     []MailjetAttachment `json:"Inline_attachments"`
	MjPrio                int                 `json:"Mj-prio"`
	MjCampaign            string              `json:"Mj-campaign"`
	MjDeduplicateCampaign bool                `json:"Mj-deduplicatecampaign"`
	MjCustomID            string              `json:"Mj-CustomID"`
	MjEventPayload        string              `json:"Mj-EventPayLoad"`
	MjTemplateID          string              `json:"Mj-Template-ID"`
	Headers               map[string]string   `json:",omitempty"`
	Vars                  interface{}
	Messages              []MailjetSendMail
}

type MailjetRecipient struct {
	Email string
	Name  string
	Vars  interface{}
}

type MailjetAttachment struct {
	ContentType string `json:"Content-Type"`
	Content     string
	Filename    string
}

type MailjetSent struct {
	Email     string
	MessageID int
}

type MailjetSentResult struct {
	Sent []MailjetSent
}
