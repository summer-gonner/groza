package registry

import (
	"strconv"
	"sync"
	"time"
)

// Registry 用于存储所有注册的服务
type Registry struct {
	services map[string]*ServiceInstance
	mu       sync.Mutex
}

// StartServer 启动服务
func (registry *Registry) StartServer(port string) {

}

// RegisterService 注册一个服务
func (registry *Registry) RegisterService(service *ServiceInstance) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	if registry.services == nil {
		registry.services = make(map[string]*ServiceInstance) // 初始化 map
	}
	serviceKey := service.ServiceName + service.Host + strconv.Itoa(service.Port)
	format := time.Now().Format("2006-01-02 15:04:05")
	service.LastHeartBeatTime = format
	service.RegisterTime = format
	registry.services[serviceKey] = service
}

// GetService 查询注册的服务
func (registry *Registry) getService(name string) *ServiceInstance {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	if service, ok := registry.services[name]; ok {
		return service
	}
	return nil
}

// GetAllServices 查询所有注册的服务
func (registry *Registry) GetAllServices() []*ServiceInstance {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	var services []*ServiceInstance
	for _, service := range registry.services {
		services = append(services, service)
	}
	return services
}

// Heartbeat 更新服务实例的心跳时间
func (registry *Registry) Heartbeat(serviceName, host string, port int) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	serviceKey := serviceName + host + string(rune(port))
	if service, exists := registry.services[serviceKey]; exists {
		service.LastHeartBeatTime = time.Now().String()
	}
}
