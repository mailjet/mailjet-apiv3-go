// Package mailjet provides methods for interacting with the last version of the Mailjet API.
// The goal of this component is to simplify the usage of the MailJet API for GO developers.
//
// For more details, see the full API Documentation at http://dev.mailjet.com/
package mailjet

import "net/http"

// ClientInterface defines all Client functions.
type ClientInterface interface {
	APIKeyPublic() string
	APIKeyPrivate() string
	Client() *http.Client
	SetClient(client *http.Client)
	List(resource string, resp interface{}, options ...RequestOptions) (count, total int, err error)
	Get(mr *Request, resp interface{}, options ...RequestOptions) error
	Post(fmr *FullRequest, resp interface{}, options ...RequestOptions) error
	Put(fmr *FullRequest, onlyFields []string, options ...RequestOptions) error
	Delete(mr *Request, options ...RequestOptions) error
	SendMail(data *InfoSendMail, options ...RequestOptions) (*SentResult, error)
	SendMailSMTP(info *InfoSMTP) (err error)
}

// ClientInterfaceV31 defines the Client functions, including SendMailV31
type ClientInterfaceV31 interface {
	ClientInterface
	SendMailV31(data *MessagesV31, options ...RequestOptions) (*ResultsV31, error)
}
