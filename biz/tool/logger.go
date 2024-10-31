package tool

import (
	"fmt"
	"github.com/fatih/color"
	"sync"
	"time"
)

const (
	LevelError = iota //写常量时，自增
	LevelWarning
	LevelInfo
	LevelDebug
)

type Logger struct {
	level int
	mu    sync.Mutex
}

var GlobalLogger *Logger
var Level = LevelDebug
var colors = map[string]func(a ...interface{}) string{
	"Warning": color.New(color.FgYellow).Add(color.Bold).SprintFunc(),
	"Panic":   color.New(color.BgRed).Add(color.Bold).SprintFunc(),
	"Err":     color.New(color.FgRed).Add(color.Bold).SprintFunc(),
	"Info":    color.New(color.FgCyan).Add(color.Bold).SprintFunc(),
	"Debug":   color.New(color.FgWhite).Add(color.Bold).SprintFunc(),
}

func (ll *Logger) Println(prefix string, msg string) {
	c := color.New()
	ll.mu.Lock()
	defer ll.mu.Unlock()
	_, _ = c.Printf(
		"%s%s %s %s\n",
		colors[prefix]("["+prefix+"]"),
		"",
		// 离谱 https://polarisxu.studygolang.com/posts/go/why-time-use-2006/
		time.Now().Format("2006-01-02 15:04:05"),
		msg,
	)
}

func (ll *Logger) Info(format string, v ...interface{}) {
	if LevelInfo > ll.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	ll.Println("Info", msg)
}

func (ll *Logger) Error(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf(format, v...)
	ll.Println("Err", msg)
}

func Log() *Logger {
	if GlobalLogger == nil {
		GlobalLogger = new(Logger)
	}
	return GlobalLogger
}
