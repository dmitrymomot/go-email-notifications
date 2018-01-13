package mailnotifier

import "strings"

// Mailer is interface for email templates
type Mailer interface {
	GetTemplate() Templater
	GetRecipient() (email string, name string)
	GetSubject() string
}

// Mail confirmation mail template structure
type Mail struct {
	Name    string
	Email   string
	Subject string
	Tpl     Templater
}

// GetTemplate returns mail html template path
func (m *Mail) GetTemplate() Templater {
	return m.Tpl
}

// GetRecipient return recipient email address and name
func (m *Mail) GetRecipient() (email string, name string) {
	email = m.Email
	if m.Name == "" {
		ec := strings.Split(email, "@")
		name = ec[0]
	}
	return
}

// GetSubject return mail subject
func (m *Mail) GetSubject() string {
	return m.Subject
}
