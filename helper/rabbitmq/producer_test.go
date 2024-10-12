package rabbitmq

import (
	"context"
	"testing"
)

func TestInitProducer(t *testing.T) {
	type args struct {
		ctx  context.Context
		conf Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitProducer(tt.args.ctx, tt.args.conf)
		})
	}
}
