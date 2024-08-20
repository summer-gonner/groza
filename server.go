package groza_server

import "fmt"

// NewRegistry 创建一个新的注册中心实例
func NewRegistry() *Registry {
	return &Registry{
		services: make(map[string]*ServiceInstance),
	}
}

// RegisterService 注册一个服务
func (r *Registry) RegisterService(service *ServiceInstance) {
	r.mu.Lock()
	defer r.mu.Unlock()
	//service.Updated = time.Now()
	r.services[service.ServiceName] = service
	fmt.Printf("注册服务名称: %s 地址加端口%s 注册时间： \n", service.ServiceName, service.Host, service.RegisterTime)
}

// GetService 查询注册的服务
func (r *Registry) GetService(name string) *ServiceInstance {
	r.mu.Lock()
	defer r.mu.Unlock()
	if service, ok := r.services[name]; ok {
		return service
	}
	return nil
}

// GetAllServices 查询所有注册的服务
func (r *Registry) GetAllServices() []*ServiceInstance {
	r.mu.Lock()
	defer r.mu.Unlock()
	var services []*ServiceInstance
	for _, service := range r.services {
		services = append(services, service)
	}
	return services
}
