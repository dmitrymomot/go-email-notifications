package mailnotifier

import (
	"reflect"
	"testing"

	gomail "gopkg.in/gomail.v2"
)

func TestSend(t *testing.T) {
	type args struct {
		mail Mailer
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
			if err := Send(tt.args.mail); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getClient(t *testing.T) {
	tests := []struct {
		name string
		want *gomail.Dialer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTemplate(t *testing.T) {
	type args struct {
		mail Mailer
	}
	tests := []struct {
		name    string
		args    args
		wantTpl string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTpl, err := parseTemplate(tt.args.mail)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTpl != tt.wantTpl {
				t.Errorf("parseTemplate() = %v, want %v", gotTpl, tt.wantTpl)
			}
		})
	}
}
