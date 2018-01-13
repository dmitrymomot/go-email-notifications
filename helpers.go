package mailnotifier

import "strings"

// SendEmailConfirmationMail helper to send an email confirmation code
func SendEmailConfirmationMail(email, name, code string) (err error) {
	if name == "" {
		ec := strings.Split(email, "@")
		name = ec[0]
	}
	err = Send(&Mail{
		Name:    name,
		Email:   email,
		Subject: "Confirm Email on " + config.ProductName,
		Tpl: &DefaultTemplate{
			Preheader: "Please, confirm your email address.",
			Intro:     []string{"Have we got the right email address to reach you on? To confirm that you can get our emails, just click the button below."},
			Outro:     []string{"If you don’t know why you got this email, please tell us straight away so we can fix this for you.", "Need help, or have questions? Just reply to this email, we'd love to help."},
			Button: &ButtonTpl{
				Title: "Confirm My Email",
				Link:  config.ConfirmEmailLink + "?email=" + email + "&code=" + code,
			},
			RemoveEmailLink: config.RemoveEmailLink + "?email=" + email + "&code=" + code,
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
	if name == "" {
		ec := strings.Split(email, "@")
		name = ec[0]
	}
	err = Send(&Mail{
		Name:    name,
		Email:   email,
		Subject: "Reset Password on " + config.ProductName,
		Tpl: &DefaultTemplate{
			Preheader: "Please, confirm your request to reset password.",
			Intro:     []string{"We got a request to reset your password. To confirm just click the button below."},
			Outro:     []string{"If you ignore this message, your password won't be changed.", "Need help, or have questions? Just reply to this email, we'd love to help."},
			Button: &ButtonTpl{
				Title: "Reset Password",
				Link:  config.ResetPasswordLink + "?email=" + email + "&code=" + code,
			},
			RemoveEmailLink: config.RemoveEmailLink + "?email=" + email + "&code=" + code,
			Product: &ProductTpl{
				Title: config.ProductName,
				Link:  config.ProductLink,
			},
		},
	})
	return
}
