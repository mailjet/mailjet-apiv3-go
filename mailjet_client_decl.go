// Package mailjet provides methods for interacting with the last version of the Mailjet API.
// The goal of this component is to simplify the usage of the MailJet API for GO developers.
//
// For more details, see the full API Documentation at http://dev.mailjet.com/
package mailjet

import "net/http"

// ClientInterface defines all Clients fuctions.
type ClientInterface interface {
	APIKeyPublic() string
	APIKeyPrivate() string
	Client() *http.Client
	SetClient(client *http.Client)
	List(resource string, resp interface{}, options ...RequestOptions) (count, total int, err error)
	Get(mr *Request, resp interface{}, options ...RequestOptions) error
	Post(fmr *FullRequest, resp interface{}, options ...RequestOptions) error
	Put(fmr *FullRequest, onlyFields []string, options ...RequestOptions) error
	Delete(mr *Request) error
	SendMail(data *InfoSendMail) (*SentResult, error)
	SendMailSMTP(info *InfoSMTP) (err error)
}
