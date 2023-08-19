package utils

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strings"
	"video_service/internal/constants"
	"video_service/internal/middleware/minio"
)

// URLconvert Convert the path in the database into a complete url accessible by the front end
func URLconvert(ctx context.Context, path string) (fullURL string) {
	if len(path) == 0 {
		return ""
	}
	arr := strings.Split(path, "/")
	hlog.Infof("arr:%+v", arr)
	u, err := minio.GetObjURL(ctx, arr[0], arr[1])
	if err != nil {
		hlog.CtxInfof(ctx, err.Error())
		return ""
	}
	// TODO ?
	u.Scheme = constants.MinioURLConvertScheme
	u.Host = constants.MinioURLConvertHost
	u.Path = "/src" + u.Path
	return u.String()
}

func NewFileName(userId, time int64) string {
	return fmt.Sprintf("%d_%d", userId, time)
}
