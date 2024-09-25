package main

import "testing"

func Test_validateEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "email", args: args{email: "user@gmail.com"}},
		{name: "email", args: args{email: "user@example.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validateEmail(tt.args.email)
		})
	}
}
