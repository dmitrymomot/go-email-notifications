package mailnotifier

import (
	"bytes"
	"html/template"
	"log"

	gomail "gopkg.in/gomail.v2"
)

// Send email notification sending function
func Send(mail Mailer) (err error) {
	tplStr, err := parseTemplate(mail)
	if err != nil {
		log.Println(err)
		return
	}

	m := gomail.NewMessage()
	m.SetAddressHeader("From", config.SenderEmail, config.getSenderName())
	email, name := mail.GetRecipient()
	m.SetAddressHeader("To", email, name)
	m.SetHeader("Subject", mail.GetSubject())
	m.SetBody("text/html", tplStr)

	client := getClient()
	if err = client.DialAndSend(m); err != nil {
		log.Println(err)
		return
	}

	return
}

func getClient() *gomail.Dialer {
	return gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
}

func parseTemplate(mail Mailer) (tpl string, err error) {
	t := template.Must(template.New("").Parse(mail.GetTemplate().HTML()))
	var buffer bytes.Buffer
	if err = t.Execute(&buffer, config); err != nil {
		log.Println(err)
		return
	}
	if err = t.Execute(&buffer, mail); err != nil {
		log.Println(err)
		return
	}
	tpl = buffer.String()
	return
}
