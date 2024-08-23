package utils

import (
	"github.com/summer-gonner/groza/logging"
	"time"
)

// CompareTime 比较时间工具函数
func CompareTime(targetTime string, duration int) bool {
	// 定义时间格式
	layout := "2006-01-02 15:04:05"
	parse, _ := time.Parse(targetTime, layout)
	currentTime := time.Now()
	subTime := currentTime.Sub(parse)
	minute := time.Minute
	durationTime := time.Duration(duration) * minute
	if subTime.Seconds() > durationTime.Seconds() {
		logging.Info("subTime", subTime.Seconds())
		logging.Info("durationTime为", durationTime.Seconds())
		return true
	} else {
		return false
	}
}
