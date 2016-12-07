package mailjet

import (
  "encoding/csv"
  "fmt"
  "net/http"
)

type httpClient struct {
  client        *http.Client
  apiKeyPublic  string
  apiKeyPrivate string
  headers       map[string]string
  request       *http.Request
  response      interface{}
}

// APIKeyPublic returns the public key.
func (c *httpClient) APIKeyPublic() string {
  return c.apiKeyPublic
}

// APIKeyPrivate returns the secret key.
func (c *httpClient) APIKeyPrivate() string {
  return c.apiKeyPrivate
}

func (c *httpClient) Client() *http.Client {
  return c.client
}

func (c *httpClient) SetClient(client *http.Client) {
  c.client = client
}

func (c *httpClient) Send(req *http.Request) *httpClient {
  c.request = req
  return c
}

func (c *httpClient) With(headers map[string]string) *httpClient {
  c.headers = headers
  return c
}

func (c *httpClient) Read(response interface{}) *httpClient {
  c.response = response
  return c
}

func (c *httpClient) Call() (count, total int, err error) {
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
      contentType := resp.Header["Content-Type"][0]
      if contentType == "application/json" {
        return readJSONResult(resp.Body, c.response)
      } else if contentType == "text/csv" {
        c.response, err = csv.NewReader(resp.Body).ReadAll()
      }
    }
  }

  return count, total, err
}

func (c *httpClient) reset() {
  c.headers = make(map[string]string)
  c.request = nil
  c.response = nil
}
