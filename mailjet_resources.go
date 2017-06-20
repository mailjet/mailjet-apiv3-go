package mailjet

import (
	"net/http"
	"net/textproto"
	"sync"

	"encoding/json"
)

/*
** API structures
 */

// Client bundles data needed by a large number
// of methods in order to interact with the Mailjet API.
type Client struct {
	apiBase    string
	httpClient HTTPClientInterface
	smtpClient SMTPClientInterface
	sync.Mutex
}

// Request bundles data needed to build the URL.
type Request struct {
	Resource string
	ID       int64
	AltID    string
	Action   string
	ActionID int64
}

// DataRequest bundles data needed to build the DATA URL.
type DataRequest struct {
	SourceType   string
	SourceTypeID int64
	DataType     string
	MimeType     string
	DataTypeID   int64
	LastID       bool
}

// FullRequest is the same as a Request but with a payload.
type FullRequest struct {
	Info    *Request
	Payload interface{}
}

// FullDataRequest is the same as a DataRequest but with a payload.
type FullDataRequest struct {
	Info    *DataRequest
	Payload interface{}
}

// RequestOptions are functional options that modify the specified request.
type RequestOptions func(*http.Request)

// RequestResult is the JSON result sent by the API.
type RequestResult struct {
	Count int
	Data  interface{}
	Total int
}

// RequestError is the error returned by the API.
type RequestError struct {
	ErrorInfo    string
	ErrorMessage string
	StatusCode   int
}

// RequestErrorV31 is the error returned by the API.
type RequestErrorV31 struct {
	ErrorInfo       string
	ErrorMessage    string
	StatusCode      int
	ErrorIdentifier string
}

/*
** Send API structures
 */

// InfoSendMail bundles data used by the Send API.
type InfoSendMail struct {
	FromEmail                string
	FromName                 string
	Sender                   string      `json:",omitempty"`
	Recipients               []Recipient `json:",omitempty"`
	To                       string      `json:",omitempty"`
	Cc                       string      `json:",omitempty"`
	Bcc                      string      `json:",omitempty"`
	Subject                  string
	TextPart                 string            `json:"Text-part,omitempty"`
	HTMLPart                 string            `json:"Html-part,omitempty"`
	Attachments              []Attachment      `json:",omitempty"`
	InlineAttachments        []Attachment      `json:"Inline_attachments,omitempty"`
	MjPrio                   int               `json:"Mj-prio,omitempty"`
	MjCampaign               string            `json:"Mj-campaign,omitempty"`
	MjDeduplicateCampaign    bool              `json:"Mj-deduplicatecampaign,omitempty"`
	MjCustomID               string            `json:"Mj-CustomID,omitempty"`
	MjTemplateID             string            `json:"Mj-TemplateID,omitempty"`
	MjTemplateErrorReporting string            `json:"MJ-TemplateErrorReporting,omitempty"`
	MjTemplateLanguage       string            `json:"Mj-TemplateLanguage,omitempty"`
	MjTemplateErrorDeliver   string            `json:"MJ-TemplateErrorDeliver,omitempty"`
	MjEventPayLoad           string            `json:"Mj-EventPayLoad,omitempty"`
	Headers                  map[string]string `json:",omitempty"`
	Vars                     interface{}       `json:",omitempty"`
	Messages                 []InfoSendMail    `json:",omitempty"`
}

// Recipient bundles data on the target of the mail.
type Recipient struct {
	Email string
	Name  string
	Vars  interface{} `json:",omitempty"`
}

// Attachment bundles data on the file attached to the mail.
type Attachment struct {
	ContentType string `json:"Content-Type"`
	Content     string
	Filename    string
}

// SentResult is the JSON result sent by the Send API.
type SentResult struct {
	Sent []struct {
		Email     string
		MessageID int64
	}
}

/*
** SMTP mail sending structures
 */

// InfoSMTP contains mandatory informations to send a mail via SMTP.
type InfoSMTP struct {
	From       string
	Recipients []string
	Header     textproto.MIMEHeader
	Content    []byte
}

/*
** Send API v3.1 structures
 */

// MessagesV31 definition
type MessagesV31 struct {
	Info        []InfoMessagesV31 `json:"Messages,omitempty"`
	SandBoxMode bool              `json:",omitempty"`
}

// InfoMessagesV31 represents the payload input taken by send API v3.1
type InfoMessagesV31 struct {
	From                     *RecipientV31          `json:",omitempty"`
	ReplyTo                  *RecipientV31          `json:",omitempty"`
	Sender                   *RecipientV31          `json:",omitempty"`
	To                       *RecipientsV31         `json:",omitempty"`
	Cc                       *RecipientsV31         `json:",omitempty"`
	Bcc                      *RecipientsV31         `json:",omitempty"`
	Attachments              *AttachmentsV31        `json:",omitempty"`
	InlinedAttachments       *InlinedAttachmentsV31 `json:",omitempty"`
	Subject                  string                 `json:",omitempty"`
	TextPart                 string                 `json:",omitempty"`
	HTMLPart                 string                 `json:",omitempty"`
	Priority                 int                    `json:",omitempty"`
	CustomCampaign           string                 `json:",omitempty"`
	StatisticsContactsListID int                    `json:",omitempty"`
	MonitoringCategory       string                 `json:",omitempty"`
	DeduplicateCampaign      bool                   `json:",omitempty"`
	TrackClicks              string                 `json:",omitempty"`
	TrackOpens               string                 `json:",omitempty"`
	CustomID                 string                 `json:",omitempty"`
	Variables                map[string]interface{} `json:",omitempty"`
	EventPayload             string                 `json:",omitempty"`
	TemplateID               interface{}            `json:",omitempty"`
	TemplateLanguage         bool                   `json:",omitempty"`
	TemplateErrorReporting   *RecipientV31          `json:",omitempty"`
	TemplateErrorDeliver     bool                   `json:",omitempty"`
	Headers                  map[string]interface{} `json:",omitempty"`
}

// RecipientV31 struct handle users input
type RecipientV31 struct {
	Email string `json:",omitempty"`
	Name  string `json:",omitempty"`
}

// RecipientsV31 is a collection of emails
type RecipientsV31 []RecipientV31

// AttachmentV31 struct represent a content attachment
type AttachmentV31 struct {
	ContentType   string `json:"ContentType,omitempty"`
	Base64Content string `json:"Base64Content,omitempty"`
	Filename      string `json:"Filename,omitempty"`
}

// AttachmentsV31 collection
type AttachmentsV31 []AttachmentV31

// InlinedAttachmentV31 struct represent the content of an inline attachement
type InlinedAttachmentV31 struct {
	AttachmentV31 `json:",omitempty"`
	ContentID     string `json:"ContentID,omitempty"`
}

// InlinedAttachmentsV31 collection
type InlinedAttachmentsV31 []InlinedAttachmentV31

// ErrorInfoV31 struct
type ErrorInfoV31 struct {
	Identifier string `json:"ErrorIdentifier,omitempty"`
	Info       string `json:"ErrorInfo"`
	Message    string `json:"ErrorMessage"`
	StatusCode int    `json:"StatusCode"`
}

func (err *ErrorInfoV31) Error() string {
	raw, _ := json.Marshal(err)
	return string(raw)
}

// APIErrorDetailsV31 contains the information details describing a specific error
type APIErrorDetailsV31 struct {
	ErrorClass     string
	ErrorMessage   string
	ErrorRelatedTo []string
	StatusCode     int
}

// APIFeedbackErrorV31 struct is composed of an error definition and the payload associated
type APIFeedbackErrorV31 struct {
	Errors []APIErrorDetailsV31
}

// APIFeedbackErrorsV31 defines the error when a validation error is being sent by the API
type APIFeedbackErrorsV31 struct {
	Messages []APIFeedbackErrorV31
}

func (api *APIFeedbackErrorsV31) Error() string {
	raw, _ := json.Marshal(api)
	return string(raw)
}

// GeneratedMessageV31 contains info to retrieve a generated email
type GeneratedMessageV31 struct {
	Email       string
	MessageUUID string
	MessageID   int64
	MessageHref string
}

// ResultV31 bundles the results of a sent email
type ResultV31 struct {
	Status   string
	CustomID string `json:",omitempty"`
	To       []GeneratedMessageV31
	Cc       []GeneratedMessageV31
	Bcc      []GeneratedMessageV31
}

// ResultsV31 bundles several results when several mails are sent
type ResultsV31 struct {
	ResultsV31 []ResultV31 `json:"Messages"`
}
