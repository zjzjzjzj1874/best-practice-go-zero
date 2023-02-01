package helper

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// GetDurationByFfmpeg ffmpeg 获取音视频时长
func GetDurationByFfmpeg(url string) (duration float64) {
	var ffmpegResp struct {
		Streams []struct {
			Duration string `json:"duration"`
		} `json:"streams"`
		Format struct {
			Duration string `json:"duration"`
		} `json:"format"`
	}

	resp, err := CommandWithTimeOut(time.Minute*2, "ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", "-i", url)
	if err != nil {
		logrus.Errorf("[GetDurationByFfmpeg]获取音视频时长失败,err:%v", err)
		return
	}
	logrus.Infof("[GetDurationByFfmpeg]获取音视频时长响应:%v", string(resp))

	_ = json.Unmarshal([]byte(resp), &ffmpegResp)
	for _, stream := range ffmpegResp.Streams {
		sduration, _ := strconv.ParseFloat(stream.Duration, 64)
		if sduration > duration {
			duration = sduration
		}
	}
	if duration == 0 {
		sduration, _ := strconv.ParseFloat(ffmpegResp.Format.Duration, 64)
		if sduration > duration {
			duration = sduration
		}
	}

	return duration
}
