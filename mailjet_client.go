// Package mailjet provides methods for interacting with the last version of the Mailjet API.
// The goal of this component is to simplify the usage of the MailJet API for GO developers.
//
// For more details, see the full API Documentation at http://dev.mailjet.com/
package mailjet

import (
	"net/http"
)

// Client bundles data needed by a large number
// of methods in order to interact with the Mailjet API.
type Client struct {
	apiBase string
	http    HTTPClientInterface
	smtp    smtpClientInterface
}

// NewMailjetClient returns a new MailjetClient using an public apikey
// and an secret apikey to be used when authenticating to API.
func NewMailjetClient(apiKeyPublic, apiKeyPrivate string, baseURL ...string) *Client {
	httpClient := NewHTTPClient(apiKeyPublic, apiKeyPrivate)
	if len(baseURL) > 0 {
		return &Client{http: httpClient, smtp: new(smtpClient), apiBase: baseURL[0]}
	}
	return &Client{http: httpClient, smtp: new(smtpClient), apiBase: apiBase}
}

// NewMailjetClientBis takes in parameter an http and smtp wrapper thus simplifying unit testing
func NewMailjetClientBis(httpClient HTTPClientInterface, smtpClient smtpClientInterface, baseURL ...string) *Client {
	if len(baseURL) > 0 {
		return &Client{http: httpClient, smtp: smtpClient, apiBase: baseURL[0]}
	}
	return &Client{http: httpClient, smtp: smtpClient, apiBase: apiBase}
}

// APIKeyPublic returns the public key.
func (c *Client) APIKeyPublic() string {
	return c.http.APIKeyPublic()
}

// APIKeyPrivate returns the secret key.
func (c *Client) APIKeyPrivate() string {
	return c.http.APIKeyPrivate()

}

// Client returns the underlying http client
func (c *Client) Client() *http.Client {
	return c.http.Client()
}

// SetClient allows to customize http client.
func (c *Client) SetClient(client *http.Client) {
	c.http.SetClient(client)
}

// List issues a GET to list the specified resource
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
func (c *Client) List(resource string, resp interface{}, options ...RequestOptions) (count, total int, err error) {
	url := buildURL(c.apiBase, &Request{Resource: resource})
	req, err := createRequest("GET", url, nil, nil, options...)
	if err != nil {
		return count, total, err
	}

	return c.http.Send(req).Read(resp).Call()
}

// Get issues a GET to view a resource specifying an id
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
// Without an specified ID in MailjetRequest, it is the same as List.
func (c *Client) Get(mr *Request, resp interface{}, options ...RequestOptions) (err error) {
	url := buildURL(c.apiBase, mr)
	req, err := createRequest("GET", url, nil, nil, options...)
	if err != nil {
		return err
	}

	_, _, err = c.http.Send(req).Read(resp).Call()
	return err
}

// Post issues a POST to create a new resource
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
func (c *Client) Post(fmr *FullRequest, resp interface{}, options ...RequestOptions) (err error) {
	url := buildURL(c.apiBase, fmr.Info)
	req, err := createRequest("POST", url, fmr.Payload, nil, options...)
	if err != nil {
		return err
	}

	headers := map[string]string{"Content-Type": "application/json"}
	_, _, err = c.http.Send(req).With(headers).Read(resp).Call()

	return err
}

// Put is used to update a resource.
// Fields to be updated must be specified by the string array onlyFields.
// If onlyFields is nil, all fields except these with the tag read_only, are updated.
// Filters can be add via functional options.
func (c *Client) Put(fmr *FullRequest, onlyFields []string, options ...RequestOptions) (err error) {
	url := buildURL(c.apiBase, fmr.Info)
	req, err := createRequest("PUT", url, fmr.Payload, onlyFields, options...)
	if err != nil {
		return err
	}

	headers := map[string]string{"Content-Type": "application/json"}
	_, _, err = c.http.Send(req).With(headers).Call()

	return err
}

// Delete is used to delete a resource.
func (c *Client) Delete(mr *Request) (err error) {
	url := buildURL(c.apiBase, mr)
	req, err := createRequest("DELETE", url, nil, nil)
	if err != nil {
		return err
	}

	_, _, err = c.http.Send(req).Call()
	return err
}

// SendMail send mail via API.
func (c *Client) SendMail(data *InfoSendMail) (res *SentResult, err error) {
	url := c.apiBase + "/send/message"
	req, err := createRequest("POST", url, data, nil)
	if err != nil {
		return res, err
	}

	headers := map[string]string{"Content-Type": "application/json"}
	_, _, err = c.http.Send(req).With(headers).Read(&res).Call()
	return res, err
}
