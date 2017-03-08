package mailjet

// NewMockedMailjetClient returns an instance of `Client` with mocked http and smtp clients injected
func NewMockedMailjetClient() *Client {
	mj := NewhttpClientMock("test", "test")

	return &Client{http: mj, smtp: new(smtpClientMock), apiBase: apiBase}
}
