package groza_server

// NewRegistryDefault 创建一个新的注册中心实例
func NewRegistryDefault() *Registry {
	return &Registry{
		services: make(map[string]*ServiceInstance),
	}
}
