// Package mailjet provides methods for interacting with the last version of the Mailjet API.
// The goal of this component is to simplify the usage of the MailJet API for GO developers.
//
// For more details, see the full API Documentation at http://dev.mailjet.com/
package mailjet

import (
  "net/http"
  "strings"
  "io"
  "os"
  "log"
)

// NewMailjetClient returns a new MailjetClient using an public apikey
// and an secret apikey to be used when authenticating to API.
func NewMailjetClient(apiKeyPublic, apiKeyPrivate string, baseUrl ...string) *Client {
  mj := &httpClient{
    apiKeyPublic: apiKeyPublic,
    apiKeyPrivate: apiKeyPrivate,
    client: http.DefaultClient,
  }
  if len(baseUrl) > 0 {
    return &Client{client: mj, apiBase:baseUrl[0]}
  }
  return &Client{client: mj, apiBase:apiBase}
}

// APIKeyPublic returns the public key.
func (c *Client) APIKeyPublic() string {
  return c.client.APIKeyPublic()
}

// APIKeyPrivate returns the secret key.
func (c *Client) APIKeyPrivate() string {
  return c.client.APIKeyPrivate()
  
}

func (c *Client) Client() *http.Client {
  return c.client.Client()
}

// SetClient allows to customize http client.
func (c *Client) SetClient(client *http.Client) {
  c.client.SetClient(client)
}

// Filter applies a filter with the defined key and value.
func Filter(key, value string) RequestOptions {
  return func(req *http.Request) {
    q := req.URL.Query()
    q.Add(key, value)
    req.URL.RawQuery = strings.Replace(q.Encode(), "%2B", "+", 1)
  }
}

// SortOrder defines the order of the result.
type SortOrder int

// These are the two possible order.
const (
  SortDesc = SortOrder(iota)
  SortAsc
)

var debugOut io.Writer = os.Stderr

// SetDebugOutput sets the output destination for the debug.
func SetDebugOutput(w io.Writer) {
  debugOut = w
  log.SetOutput(w)
}

// Sort applies the Sort filter to the request.
func Sort(value string, order SortOrder) RequestOptions {
  if order == SortDesc {
    value = value + "+DESC"
  }
  return Filter("Sort", value)
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

  return c.client.Send(req).Read(resp).Call()
}

// Get issues a GET to view a resource specifying an id
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
// Without an specified ID in MailjetRequest, it is the same as List.
func (mj *Client) Get(mr *Request, resp interface{}, options ...RequestOptions) (err error) {
  url := buildURL(mj.apiBase, mr)
  req, err := createRequest("GET", url, nil, nil, options...)
  if err != nil {
    return err
  }

  _, _, err = mj.client.Send(req).Read(resp).Call()
  return err
}

// Post issues a POST to create a new resource
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
func (mj *Client) Post(fmr *FullRequest, resp interface{}, options ...RequestOptions) (err error) {
  url := buildURL(mj.apiBase, fmr.Info)
  req, err := createRequest("POST", url, fmr.Payload, nil, options...)
  if err != nil {
    return err
  }

  headers := map[string]string{"Content-Type": "application/json"}
  _,_, err = mj.client.Send(req).With(headers).Read(resp).Call()

  return err
}

// Put is used to update a resource.
// Fields to be updated must be specified by the string array onlyFields.
// If onlyFields is nil, all fields except these with the tag read_only, are updated.
// Filters can be add via functional options.
func (mj *Client) Put(fmr *FullRequest, onlyFields []string, options ...RequestOptions) (err error) {
  url := buildURL(mj.apiBase, fmr.Info)
  req, err := createRequest("PUT", url, fmr.Payload, onlyFields, options...)
  if err != nil {
    return err
  }

  headers := map[string]string{"Content-Type": "application/json"}
  _, _, err = mj.client.Send(req).With(headers).Call()

  return err
}

// Delete is used to delete a resource.
func (mj *Client) Delete(mr *Request) (err error) {
  url := buildURL(mj.apiBase, mr)
  req, err := createRequest("DELETE", url, nil, nil)
  if err != nil {
    return err
  }

  _, _, err = mj.client.Send(req).Call()
  return err
}

// SendMail send mail via API.
func (mj *Client) SendMail(data *InfoSendMail) (res *SentResult, err error) {
  url := mj.apiBase + "/send/message"
  req, err := createRequest("POST", url, data, nil)
  if err != nil {
    return res, err
  }

  headers := map[string]string{"Content-Type": "application/json"}
  
  _, _, err = mj.client.Send(req).With(headers).Read(&res).Call()

  return res, err
}
