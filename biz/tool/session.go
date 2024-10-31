package tool

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// SetSession 设置session
func SetSession(c *app.RequestContext, list map[string]interface{}) {
	s := sessions.Default(c)
	for key, value := range list {
		s.Set(key, value)
	}

	err := s.Save()
	if err != nil {
		Log().Error("无法设置 Session 值：%s", err)
	}
}

// GetSession 获取session
func GetSession(c *app.RequestContext, key string) interface{} {
	s := sessions.Default(c)
	return s.Get(key)
}

// DeleteSession 删除session
func DeleteSession(c *app.RequestContext, key string) {
	s := sessions.Default(c)
	s.Delete(key)
	s.Save()
}

// ClearSession 清空session
func ClearSession(c *app.RequestContext) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
}
