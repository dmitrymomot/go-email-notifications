package mailnotifier

import (
	"os"
	"testing"
)

func TestSetupConfig(t *testing.T) {
	type args struct {
		c *Config
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Setup Config 1", args{&Config{SenderEmail: "test@mail"}}},
		{"Test Setup Config 2", args{&Config{SenderName: "Sender Name"}}},
		{"Test Setup Config 3", args{&Config{ProductName: "App", SenderName: "Sender Name"}}},
		{"Test Setup Config 4", args{&Config{ProductName: "App"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetupConfig(tt.args.c)
			if res := config.getSenderName(); res != tt.args.c.getSenderName() {
				t.Errorf("getSenderName() returns unexpected value %v", res)
			}
		})
	}
}

func TestInitFromDotenvFile(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Test Load Config From Dotenv File: success", false},
		{"Test Load Config From Dotenv File: failed", true},
	}
	for _, tt := range tests {
		_, exists := os.Stat(".env")
		if tt.wantErr == true && exists == nil {
			os.Rename(".env", ".env.sample")
		}
		if tt.wantErr == false && exists != nil {
			os.Rename(".env.sample", ".env")
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := InitFromDotenvFile(); (err != nil) != tt.wantErr {
				t.Errorf("InitFromDotenvFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
