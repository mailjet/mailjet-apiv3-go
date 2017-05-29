package mailjet

import "net/http"

// HTTPClientInterface method definition
type HTTPClientInterface interface {
	APIKeyPublic() string
	APIKeyPrivate() string
	Client() *http.Client
	SetClient(client *http.Client)
	Send(req *http.Request) HTTPClientInterface
	SendMailV31(req *http.Request) (*http.Response, error)
	With(headers map[string]string) HTTPClientInterface
	Read(response interface{}) HTTPClientInterface
	Call() (count int, total int, err error)
}
