package helper

import "testing"

func TestGetDurationByFfmpeg(t *testing.T) {
	duration := GetDurationByFfmpeg("https://www.baidu.com/2.mp4")
	t.Log(duration)
}
