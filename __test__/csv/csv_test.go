package main

import "testing"

func Test_readCsv(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "readCsv", args: args{filename: "/Users/zjzjzjzj1874/personal/best-practice-go-zero/go.csv"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readCsv(tt.args.filename)
		})
	}
}
