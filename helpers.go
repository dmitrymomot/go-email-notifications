package mailnotifier

// SendEmailConfirmationMail helper to send an email confirmation code
func SendEmailConfirmationMail(email, name, code string) (err error) {
	err = Send(&Mail{
		Name:    name,
		Email:   email,
		Subject: "Confirm Email",
		Tpl: &DefaultTemplate{
			Preheader: "preheader test string",
			Intro:     []string{"test intro"},
			Outro:     []string{"test outro"},
			Button: &ButtonTpl{
				Title: "Confirm Email",
				Link:  config.ConfirmEmailLink,
			},
			CompanyAddress:  "Company Address test",
			UnsubscribeLink: config.UnsubscribeLink,
			RemoveEmailLink: config.RemoveEmailLink,
			Product: &ProductTpl{
				Title: config.ProductName,
				Link:  config.ProductLink,
			},
		},
	})
	return
}

// SendResetPasswordMail helper to send an email with a confirmation code to reset password
func SendResetPasswordMail(email, name, code string) (err error) {
	err = Send(&Mail{
		Name:    name,
		Email:   email,
		Subject: "Confirm Email",
		Tpl: &DefaultTemplate{
			Preheader: "preheader test string",
			Intro:     []string{"test intro"},
			Outro:     []string{"test outro"},
			Button: &ButtonTpl{
				Title: "Confirm Email",
				Link:  "http://localhost:8080",
			},
			CompanyAddress:  "Company Address test",
			UnsubscribeLink: "http://localhost:8080/unsubscribe",
			RemoveEmailLink: "http://localhost:8080/remove",
			Product: &ProductTpl{
				Title: "Pixamazer",
				Link:  "http://pixamazer.com",
			},
		},
	})
	return
}
