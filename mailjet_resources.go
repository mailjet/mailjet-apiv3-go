package mailjet

import "net/http"

/*
** API structures
 */

// Client bundles data needed by a large number
// of methods in order to interact with the Mailjet API.
type Client struct {
  apiBase       string
  client        *httpClient
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
