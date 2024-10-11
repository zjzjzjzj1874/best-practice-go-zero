package gbk

import "testing"

func Test_encodeSimpleChinese(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Encode", args: args{name: "锟斤拷"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodeSimpleChinese(tt.args.name)
		})
	}
}
