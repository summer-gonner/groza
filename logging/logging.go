package logging

import (
	"github.com/sirupsen/logrus"

	"os"
)

// Logger 全局日志实例
var Logger *logrus.Logger

// TraceIDHook 是一个 logrus Hook 实现
type TraceIDHook struct {
	TraceID string
}

func (hook *TraceIDHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *TraceIDHook) Fire(entry *logrus.Entry) error {
	entry.Data["traceId"] = hook.TraceID
	return nil
}

// Init 日志初始化
func Init(traceId string) {
	// 初始化全局日志实例
	Logger = logrus.New()

	// 设置日志输出到标准输出
	Logger.SetOutput(os.Stdout)

	// 设置日志级别
	Logger.SetLevel(logrus.InfoLevel)

	// 可选：设置日志格式
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	// 返回一个带有 traceId 的日志记录器实例
	// 添加 hook
	Logger.AddHook(&TraceIDHook{TraceID: traceId})
}
