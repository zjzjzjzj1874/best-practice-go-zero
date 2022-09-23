// Package obs huawei object storage service
package obs

import (
	"os"
	"testing"

	"github.com/zjzjzjzj1874/huaweicloud-sdk-go-obs/obs"
)

var client = NewHWObsClient(ConfObs{
	AK:       "your access key",
	SK:       "your secret key",
	Endpoint: "your endpoint",
})

func TestOBS_Func(t *testing.T) {
	t.Run("upload", func(t *testing.T) {
		file, err := os.Open("./hw-obs.go")
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = file.Close() }()

		obsInput := &obs.PutObjectInput{}
		obsInput.Body = file
		obsInput.ContentDisposition = "attachment; filename=obs.go"
		obsInput.Bucket = "bucket"
		obsInput.Key = "123/dev/1234234321" // 这里拼接key，即指定存储文件的位置
		output, err := client.PutObject(obsInput)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(output.RequestId)
	})

	t.Run("upload && update metadata", func(t *testing.T) {
		file, err := os.Open("./hw-obs.go")
		if err != nil {
			t.Fatal(err)
		}
		defer func() { _ = file.Close() }()

		obsInput := &obs.PutObjectInput{}
		obsInput.Body = file
		obsInput.Bucket = "bucket"
		obsInput.Key = "123/dev/1234234321" // 这里拼接key，即指定存储文件的位置

		output, err := client.PutObject(obsInput)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(output.RequestId)

		pfi := &obs.SetObjectMetadataInput{}
		pfi.Bucket = "bucket"
		pfi.Key = "123/dev/1234234321"
		pfi.ContentDisposition = "attachment; filename=obs.go"
		out, err := client.SetObjectMetadata(pfi)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(out)

	})

	t.Run("upload", func(t *testing.T) {
		output, err := client.PutFile(&obs.PutFileInput{
			PutObjectBasicInput: obs.PutObjectBasicInput{
				ObjectOperationInput: obs.ObjectOperationInput{
					Bucket: "bucket",
					Key:    "my-zero/dev/12342343223",
				},
			},
			SourceFile: "./hw-obs.go",
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(output.RequestId)
	})

	t.Run("download", func(t *testing.T) {
		input := &obs.GetObjectInput{}
		input.Bucket = "bucket"
		input.Key = "my-zero/dev/12342343223"

		output, err := client.GetObject(input)
		if err == nil {
			defer output.Body.Close()
			t.Logf("StorageClass:%s, ETag:%s, ContentType:%s, ContentLength:%d, LastModified:%s\n",
				output.StorageClass, output.ETag, output.ContentType, output.ContentLength, output.LastModified)
		} else if obsError, ok := err.(obs.ObsError); ok {
			t.Logf("Code:%s\n", obsError.Code)
			t.Logf("Message:%s\n", obsError.Message)
		}
	})

	t.Run("signurl", func(t *testing.T) {
		input := &obs.CreateSignedUrlInput{}
		input.Bucket = "123"
		input.Key = "my_zero/dev/1234234321"
		input.Method = obs.HttpMethodGet
		input.Expires = 86400 // 图片URL超时时间

		output, err := client.CreateSignedUrl(input)
		if err == nil {
			t.Logf("url:%s.", output.SignedUrl)
		} else if obsError, ok := err.(obs.ObsError); ok {
			t.Logf("Code:%s\n", obsError.Code)
			t.Logf("Message:%s\n", obsError.Message)
		}
	})

}
