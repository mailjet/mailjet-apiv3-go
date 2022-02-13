// Package mailjet provides methods for interacting with the last version of the Mailjet API.
// The goal of this component is to simplify the usage of the MailJet API for GO developers.
//
// For more details, see the full API Documentation at http://dev.mailjet.com/
package mailjet

import (
	"context"
	"net/http"
)

// ClientInterface defines all Clients functions.
type ClientInterface interface {
	APIKeyPublic() string
	APIKeyPrivate() string
	Client() *http.Client
	SetClient(client *http.Client)
	List(resource string, resp interface{}, options ...RequestOptions) (count, total int, err error)
	ListWithContext(ctx context.Context, resource string, resp interface{}, options ...RequestOptions) (count, total int, err error)
	Get(mr *Request, resp interface{}, options ...RequestOptions) error
	GetWithContext(ctx context.Context, mr *Request, resp interface{}, options ...RequestOptions) error
	Post(fmr *FullRequest, resp interface{}, options ...RequestOptions) error
	PostWithContext(ctx context.Context, fmr *FullRequest, resp interface{}, options ...RequestOptions) error
	Put(fmr *FullRequest, onlyFields []string, options ...RequestOptions) error
	PutWithContext(ctx context.Context, fmr *FullRequest, onlyFields []string, options ...RequestOptions) error
	Delete(mr *Request) error
	DeleteWithContext(ctx context.Context, mr *Request) error
	SendMail(data *InfoSendMail) (*SentResult, error)
	SendMailWithContext(ctx context.Context, data *InfoSendMail) (*SentResult, error)
	SendMailSMTP(info *InfoSMTP) (err error)
}
