package mailjet

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

// HTTPClient is a wrapper arround http.Client
type HTTPClient struct {
	client        *http.Client
	apiKeyPublic  string
	apiKeyPrivate string
	headers       map[string]string
	request       *http.Request
	response      interface{}
}

// NewHTTPClient returns a new httpClient
func NewHTTPClient(apiKeyPublic, apiKeyPrivate string) *HTTPClient {
	return &HTTPClient{
		apiKeyPublic:  apiKeyPublic,
		apiKeyPrivate: apiKeyPrivate,
		client:        http.DefaultClient,
	}
}

// APIKeyPublic returns the public key.
func (c *HTTPClient) APIKeyPublic() string {
	return c.apiKeyPublic
}

// APIKeyPrivate returns the secret key.
func (c *HTTPClient) APIKeyPrivate() string {
	return c.apiKeyPrivate
}

// Client returns the underlying http client
func (c *HTTPClient) Client() *http.Client {
	return c.client
}

// SetClient sets the underlying http client
func (c *HTTPClient) SetClient(client *http.Client) {
	c.client = client
}

// Send binds the request to the underlying http client
func (c *HTTPClient) Send(req *http.Request) HTTPClientInterface {
	c.request = req
	return c
}

// With binds the header to the underlying http client
func (c *HTTPClient) With(headers map[string]string) HTTPClientInterface {
	c.headers = headers
	return c
}

// SendMailV31 simply calls the underlying http client.Do function
func (c *HTTPClient) SendMailV31(req *http.Request) (*http.Response, error) {
	res, err := c.Client().Do(req)
	return res, err
}

// Read binds the response to the underlying http client
func (c *HTTPClient) Read(response interface{}) HTTPClientInterface {
	c.response = response
	return c
}

// Call execute the HTTP call to the API
func (c *HTTPClient) Call() (count, total int, err error) {
	defer c.reset()
	for key, value := range c.headers {
		c.request.Header.Add(key, value)
	}

	resp, err := c.doRequest(c.request)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return count, total, err
	} else if resp == nil {
		return count, total, fmt.Errorf("empty response")
	}

	if c.response != nil {
		if resp.Header["Content-Type"] != nil {
			contentType := strings.ToLower(resp.Header["Content-Type"][0])
			if strings.Contains(contentType, "application/json") {
				return readJSONResult(resp.Body, c.response)
			} else if strings.Contains(contentType, "text/csv") {
				c.response, err = csv.NewReader(resp.Body).ReadAll()
			}
		}
	}

	return count, total, err
}

func (c *HTTPClient) reset() {
	c.headers = make(map[string]string)
	c.request = nil
	c.response = nil
}
