package minio

import (
	"context"
	"testing"
)

var (
	client *Client
	ctx    = context.Background()
	conf   = Conf{
		Endpoint: "127.0.0.1:9000",
		AK:       "minio",
		SK:       "minio123",
		SSL:      false,
		Bucket:   "testbucket",
		Location: "local",
	}
)

func init() {
	client = MustNewMinioClient(context.Background(), conf)
}
func TestMustNewMinioClient(t *testing.T) {
	type args struct {
		ctx  context.Context
		conf Conf
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "mustNewMinioClient", args: args{
			ctx:  context.Background(),
			conf: conf,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client = MustNewMinioClient(tt.args.ctx, tt.args.conf)
		})
	}

	t.Run("PutObject", func(t *testing.T) {
		client.PutObjectByPath(ctx, "testFile1", "/tmp/test.txt")
	})

	t.Run("GetObj", func(t *testing.T) {
		client.GetObject(ctx, "testFile1")
	})

	t.Run("PresignedGetObject", func(t *testing.T) {
		client.PresignedGetObject(ctx, "testFile1")
	})
}
