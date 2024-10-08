package minio

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	ctx  context.Context
	C    *minio.Client
	conf Conf
}
type Conf struct {
	Endpoint string `json:"endpoint"`
	AK       string `json:"ak"`
	SK       string `json:"sk"`
	SSL      bool   `json:"ssl"`
	Bucket   string `json:"bucket"`
	Location string `json:"location"`
}

func MustNewMinioClient(ctx context.Context, conf Conf) *Client {
	client, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AK, conf.SK, ""),
		Secure: conf.SSL,
	})
	if err != nil {
		panic(err)
	}

	exist, err := client.BucketExists(ctx, conf.Bucket)
	if err != nil {
		panic(err)
	}
	if !exist {
		err = client.MakeBucket(ctx, conf.Bucket, minio.MakeBucketOptions{
			Region: conf.Location,
		})
		if err != nil {
			panic(err)
		}
	}

	return &Client{
		ctx:  ctx,
		C:    client,
		conf: conf,
	}
}

func (c *Client) PutObjectByPath(ctx context.Context, objName, path string) {
	info, err := c.C.FPutObject(ctx, c.conf.Bucket, objName, path, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return
	}

	fmt.Println(info.Key)
}

func (c *Client) GetObject(ctx context.Context, objName string) {
	obj, err := c.C.GetObject(ctx, c.conf.Bucket, objName, minio.GetObjectOptions{})
	if err != nil {
		return
	}
	buf, err := io.ReadAll(obj)
	fmt.Println(string(buf))

}
func (c *Client) PresignedGetObject(ctx context.Context, objName string) {
	url, err := c.C.PresignedGetObject(ctx, c.conf.Bucket, objName, time.Hour*24, url.Values{})
	if err != nil {
		return
	}
	fmt.Println("url == ", url.String())
}
