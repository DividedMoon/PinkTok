package ffmpeg

import (
	"bytes"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// GetSnapshot 通过ffmpeg获取视频的第一帧
func GetSnapshot(videoPath string) (buf *bytes.Buffer, err error) {
	buf = bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf).
		Run()
	if err != nil {
		hlog.Error(err)
		return nil, err
	}

	return buf, nil
}
