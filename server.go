package groza

import (
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/groza/common/response"
	"github.com/summer-gonner/groza/logging"
	"time"
)

type Server struct {
}

// 服务注册
func (s *Server) GetService(c *gin.Context) {
	name := c.Param("name")
	response.Success(c, name)
}

func (s *Server) GetAllServices(c *gin.Context) {
	registry := Registry{}
	services := registry.getAllServices()
	logging.Info("查询到的服务信息为 %s", services)
	response.Success(c, services)
}

func (s *Server) ServiceRegister(c *gin.Context) {

	var service ServiceInstance
	err := c.ShouldBindJSON(&service)
	logging.Info(" 【sevice】 %v", service)
	registry := Registry{}
	registry.registerService(&service)
	if err != nil {
		logging.Errorf(" 【err】 %v", err)
		response.Fail(c, "注册失败")
	} else {
		response.Success(c, service)
	}

}

// HeartBeat 心跳机制
func (s *Server) HeartBeat(c *gin.Context) {

}

// StartServer 启动服务
func (registry *Registry) StartServer(port string) {

}

// RegisterService 注册一个服务
func (registry *Registry) registerService(service *ServiceInstance) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	serviceKey := service.ServiceName + service.Host + string(rune(service.Port))
	service.LastHeartBeatTime = time.Now().String()
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
func (registry *Registry) getAllServices() []*ServiceInstance {
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
