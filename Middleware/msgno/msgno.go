package msgno

import (
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"math/rand"
	"strings"
	"time"
)

func MsgNoMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		msgno, ok := metainfo.GetValue(ctx, "msgno")
		if !ok {
			msgno = generateMsgNo()
			ctx = metainfo.WithValue(ctx, "msgno", msgno)
		}
		return next(ctx, request, response)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func generateMsgNo() string {
	// 生成当前日期的字符串，格式为YYYYMMDD
	now := time.Now()
	date := now.Format("20060102")
	sb := strings.Builder{}
	sb.Grow(12)
	sb.WriteString(date)
	// 生成四位随机字符
	n := 4
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()

}
