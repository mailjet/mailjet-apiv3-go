// Package mailjet provides methods for interacting with the last version of the Mailjet API.
// The goal of this component is to simplify the usage of the MailJet API for GO developers.
//
// For more details, see the full API Documentation at http://dev.mailjet.com/
package mailjet

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// NewMailjetClient returns a new MailjetClient using an public apikey
// and an secret apikey to be used when authenticating to API.
func NewMailjetClient(apiKeyPublic, apiKeyPrivate string) *MailjetClient {
	mj := &MailjetClient{
		apiKeyPublic:  apiKeyPublic,
		apiKeyPrivate: apiKeyPrivate,
		client:        http.DefaultClient,
	}
	return mj
}

// APIKeyPublic returns the public key.
func (mj *MailjetClient) APIKeyPublic() string {
	return mj.apiKeyPublic
}

// APIKeyPrivate returns the secret key.
func (mj *MailjetClient) APIKeyPrivate() string {
	return mj.apiKeyPrivate
}

// Client returns the http client used by the wrapper.
func (mj *MailjetClient) Client() *http.Client {
	return mj.client
}

// SetClient allows to customize http client.
func (mj *MailjetClient) SetClient(c *http.Client) {
	mj.client = c
}

// SetDebugOutput sets the output destination for the debug.
func SetDebugOutput(w io.Writer) {
	debugOut = w
	log.SetOutput(w)
}

func Filter(key, value string) MailjetOptions {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add(key, value)
		req.URL.RawQuery = strings.Replace(q.Encode(), "%2B", "+", 1)
	}
}

type SortOrder int

const (
	SortDesc = SortOrder(iota)
	SortAsc
)

func Sort(value string, order SortOrder) MailjetOptions {
	if order == SortDesc {
		value = value + "+DESC"
	}
	return Filter("Sort", value)
}

// List issues a GET to list the specified resource
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
func (mj *MailjetClient) List(resource string, res interface{}, options ...MailjetOptions) (count, total int, err error) {
	url := buildURL(&MailjetRequest{Resource: resource})
	req, err := createRequest("GET", url, nil, nil, options...)
	if err != nil {
		return count, total, err
	}
	resp, err := mj.doRequest(req)
	if err != nil {
		return count, total, err
	} else if resp == nil {
		return count, total, fmt.Errorf("empty response")
	}
	defer resp.Body.Close()

	return readJSONResult(resp.Body, res)
}

// Get issues a GET to view a resource specifying an id
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
// Without an specified ID in MailjetRequest, it is the same as List.
func (mj *MailjetClient) Get(mr *MailjetRequest, res interface{}, options ...MailjetOptions) (err error) {
	url := buildURL(mr)
	req, err := createRequest("GET", url, nil, nil, options...)
	if err != nil {
		return err
	}
	resp, err := mj.doRequest(req)
	if err != nil {
		return err
	} else if resp == nil {
		return fmt.Errorf("empty response")
	}
	defer resp.Body.Close()

	_, _, err = readJSONResult(resp.Body, res)
	return err
}

// Post issues a POST to create a new resource
// and stores the result in the value pointed to by res.
// Filters can be add via functional options.
func (mj *MailjetClient) Post(fmr *FullMailjetRequest, res interface{}, options ...MailjetOptions) (err error) {
	url := buildURL(fmr.Info)
	req, err := createRequest("POST", url, fmr.Payload, nil, options...)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := mj.doRequest(req)
	if err != nil {
		return err
	} else if resp == nil {
		return fmt.Errorf("empty response")
	}
	defer resp.Body.Close()

	_, _, err = readJSONResult(resp.Body, res)
	return err
}

// Put is used to update a resource.
// Fields to be updated must be specified by the string array onlyFields.
// If onlyFields is nil, all fields except these with the tag read_only, are updated.
// Filters can be add via functional options.
func (mj *MailjetClient) Put(fmr *FullMailjetRequest, onlyFields []string, options ...MailjetOptions) (err error) {
	url := buildURL(fmr.Info)
	req, err := createRequest("PUT", url, fmr.Payload, onlyFields, options...)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := mj.doRequest(req)
	if resp != nil {
		resp.Body.Close()
	}

	return err
}

// Delete is used to delete a resource.
func (mj *MailjetClient) Delete(mr *MailjetRequest) (err error) {
	url := buildURL(mr)
	r, err := createRequest("DELETE", url, nil, nil)
	if err != nil {
		return err
	}
	resp, err := mj.doRequest(r)
	if resp != nil {
		resp.Body.Close()
	}

	return err
}

// SendMail send mail via API.
func (mj *MailjetClient) SendMail(data *MailjetSendMail) (res *MailjetSentResult, err error) {
	url := apiBase + "/send/message"
	req, err := createRequest("POST", url, data, nil)
	if err != nil {
		return res, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := mj.doRequest(req)
	if err != nil {
		return res, err
	} else if resp == nil {
		return res, fmt.Errorf("empty response")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	return res, err
}
