package middleware

import (
	"context"
	"github.com/CloudGoSight/cloudgosight_api/biz/model"
	"github.com/CloudGoSight/cloudgosight_api/conf"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"net/http"
	"strings"
)

func CurrentUser() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetActiveUserByID(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next(ctx)
	}
}

// todo 移到rpc服務
// 初始化session
// return func(ctx context.Context, c *app.RequestContext) {} 内部的逻辑在每次请求时都执行，但外部逻辑只执行一次
func Session(secret string) app.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(secret))

	sameSiteMode := http.SameSiteDefaultMode
	switch strings.ToLower(conf.CORSConfig.SameSite) {
	case "default":
		sameSiteMode = http.SameSiteDefaultMode
	case "none":
		sameSiteMode = http.SameSiteNoneMode
	case "strict":
		sameSiteMode = http.SameSiteStrictMode
	case "lax":
		sameSiteMode = http.SameSiteLaxMode
	}

	// Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{
		HttpOnly: true,
		MaxAge:   60 * 86400,
		Path:     "/",
		SameSite: sameSiteMode,
		Secure:   conf.CORSConfig.Secure,
	})

	return sessions.New("cloudgosight-session", store)
}

func CacheControl() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Header("Cache-Control", "private, no-cache")
	}
}
