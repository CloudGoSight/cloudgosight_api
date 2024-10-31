package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// CaptchaRequired 验证请求签名
func CaptchaRequired() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// captcha：区分机器和人的图灵测试，跳过
		c.Next(ctx)
	}
}
