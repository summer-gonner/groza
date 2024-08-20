package groza_server

import "time"

// Cleanup 定期清理超时的服务实例
func (r *Registry) Cleanup(timeout time.Duration) {
	for {
		time.Sleep(timeout)

		r.mu.Lock()
		for key, service := range r.services {
			// 定义时间格式
			layout := "2006-01-02 15:04:05"
			parse, _ := time.Parse(layout, service.LastHeartBeatTime)
			if time.Since(parse) > timeout {
				delete(r.services, key)
			}
		}
		r.mu.Unlock()
	}
}

// Cleanup 定期清理超时的服务实例
func (r *Registry) CleanupDefault() {
	for {
		timeout := 15 * time.Second
		time.Sleep(timeout)

		r.mu.Lock()
		for key, service := range r.services {
			// 定义时间格式
			layout := "2006-01-02 15:04:05"
			parse, _ := time.Parse(layout, service.LastHeartBeatTime)
			if time.Since(parse) > timeout {
				delete(r.services, key)
			}
		}
		r.mu.Unlock()
	}
}
