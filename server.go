package groza_server

import "fmt"

// RegisterService 注册一个服务
func (registry *Registry) RegisterService(service *ServiceInstance) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	registry.services[service.ServiceName] = service
	fmt.Printf("注册服务名称: %s 地址加端口%s 注册时间：%s", service.ServiceName, service.Host, service.RegisterTime)
}

// GetService 查询注册的服务
func (registry *Registry) GetService(name string) *ServiceInstance {
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
