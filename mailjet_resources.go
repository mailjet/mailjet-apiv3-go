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
	Sender                string             `json:",omitempty"`
	Recipients            []MailjetRecipient `json:",omitempty"`
	To                    []string           `json:",omitempty"`
	Cc                    []string           `json:",omitempty"`
	Bcc                   []string           `json:",omitempty"`
	Subject               string
	TextPart              string              `json:"Text-part,omitempty"`
	HtmlPart              string              `json:"Html-part,omitempty"`
	Attachments           []MailjetAttachment `json:",omitempty"`
	InlineAttachments     []MailjetAttachment `json:"Inline_attachments,omitempty"`
	MjPrio                int                 `json:"Mj-prio,omitempty"`
	MjCampaign            string              `json:"Mj-campaign,omitempty"`
	MjDeduplicateCampaign bool                `json:"Mj-deduplicatecampaign,omitempty"`
	MjCustomID            string              `json:"Mj-CustomID,omitempty"`
	MjEventPayLoad        string              `json:"Mj-EventPayLoad,omitempty"`
	Headers               map[string]string   `json:",omitempty"`
	Vars                  interface{}         `json:",omitempty"`
	Messages              []MailjetSendMail   `json:",omitempty"`
}

type MailjetRecipient struct {
	Email string
	Name  string
	Vars  interface{} `json:",omitempty"`
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
