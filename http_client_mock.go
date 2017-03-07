package mailjet

import (
	"errors"
	"net/http"

	"github.com/mailjet/mailjet-apiv3-go/fixtures"
)

// HTTPClientMock definition
type HTTPClientMock struct {
	client        *http.Client
	apiKeyPublic  string
	apiKeyPrivate string
	headers       map[string]string
	request       *http.Request
	response      interface{}
	validCreds    bool
	f             *fixtures.Fixtures
}

// NewhttpClientMock instanciate new httpClientMock
func NewhttpClientMock(apiKeyPublic, apiKeyPrivate string) *HTTPClientMock {

	return &HTTPClientMock{
		apiKeyPublic:  apiKeyPublic,
		apiKeyPrivate: apiKeyPrivate,
		client:        http.DefaultClient,
		validCreds:    true,
		f:             fixtures.New(),
	}
}

// APIKeyPublic returns the public key.
func (c *HTTPClientMock) APIKeyPublic() string {
	return c.apiKeyPublic
}

// APIKeyPrivate returns the secret key.
func (c *HTTPClientMock) APIKeyPrivate() string {
	return c.apiKeyPrivate
}

// Client returns the underlying http client
func (c *HTTPClientMock) Client() *http.Client {
	return c.client
}

// SetClient allow to set the underlying http client
func (c *HTTPClientMock) SetClient(client *http.Client) {
	c.client = client
}

// Send data through HTTP with the current configuration
func (c *HTTPClientMock) Send(req *http.Request) HTTPClientInterface {
	c.request = req
	return c
}

// With lets you set the http header and returns the HTTPClientMock with the header modified
func (c *HTTPClientMock) With(headers map[string]string) HTTPClientInterface {
	c.headers = headers
	return c
}

// Read allow you to bind the response recieved through the underlying http client
func (c *HTTPClientMock) Read(response interface{}) HTTPClientInterface {
	c.f.Read(response)
	return c
}

// Call the mailjet API
func (c *HTTPClientMock) Call() (int, int, error) {
	if c.validCreds == true {
		return 1, 1, nil
	}
	return 0, 0, errors.New("Unexpected error: Unexpected server response code: 401: EOF")
}
