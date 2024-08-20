package groza_server

// NewRegistry 创建一个新的注册中心实例 需要配置启动端口
func NewRegistry(port string) *Registry {
	return &Registry{
		services: make(map[string]*ServiceInstance),
		Port:     port,
	}
}

// NewRegistryDefault 创建一个新的注册中心实例 默认给一个端口 8848
func NewRegistryDefault() *Registry {
	return &Registry{
		services: make(map[string]*ServiceInstance),
		Port:     "8848",
	}
}
