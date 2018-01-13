package mailnotifier

import "testing"

func TestSendEmailConfirmationMail(t *testing.T) {
	type args struct {
		email string
		name  string
		code  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendEmailConfirmationMail(tt.args.email, tt.args.name, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("SendEmailConfirmationMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendResetPasswordMail(t *testing.T) {
	type args struct {
		email string
		name  string
		code  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendResetPasswordMail(tt.args.email, tt.args.name, tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("SendResetPasswordMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
