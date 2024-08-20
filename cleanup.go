package groza

import "time"

// Cleanup 定期清理超时的服务实例
func (registry *Registry) Cleanup(timeout time.Duration) {
	for {
		time.Sleep(timeout)

		registry.mu.Lock()
		for key, service := range registry.services {
			// 定义时间格式
			layout := "2006-01-02 15:04:05"
			parse, _ := time.Parse(layout, service.LastHeartBeatTime)
			if time.Since(parse) > timeout {
				delete(registry.services, key)
			}
		}
		registry.mu.Unlock()
	}
}

// CleanupDefault 定期清理超时的服务实例
func (registry *Registry) CleanupDefault() {
	for {
		timeout := 15 * time.Second
		time.Sleep(timeout)

		registry.mu.Lock()
		for key, service := range registry.services {
			// 定义时间格式
			layout := "2006-01-02 15:04:05"
			parse, _ := time.Parse(layout, service.LastHeartBeatTime)
			if time.Since(parse) > timeout {
				delete(registry.services, key)
			}
		}
		registry.mu.Unlock()
	}
}
