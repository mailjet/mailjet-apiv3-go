package mailjet

// SMTPClientInterface def
type SMTPClientInterface interface {
	SendMail(from string, to []string, msg []byte) error
}
