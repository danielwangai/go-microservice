package mailer

import "net/smtp"

type MailConfig struct {
	mailServerAddress string
	auth              smtp.Auth
	from              string
	username          string
	password          string
	host              string
}

func NewMailAuth(username, password, host string) smtp.Auth {
	return smtp.PlainAuth("", username, password, host)
}

func NewMailConfig(mailServerAddress, from, username, password, host string, auth smtp.Auth) *MailConfig {
	return &MailConfig{mailServerAddress, auth, from, username, password, host}
}

func (m *MailConfig) SendMail(to []string, content string) error {
	auth := smtp.PlainAuth("", m.from, m.password, m.host)
	err := smtp.SendMail(m.mailServerAddress, auth, m.from, to, []byte(content))
	if err != nil {
		return err
	}

	return nil
}
