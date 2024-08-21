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
func (s *Server) getService(c *gin.Context) {
	name := c.Param("name")
	response.Success(c, name)
}

func (s *Server) getAllServices(c *gin.Context) {

}

func (s *Server) serviceRegister(c *gin.Context) {

	var service ServiceInstance
	err := c.ShouldBindJSON(&service)
	logging.Logger.Printf(" 【sevice】 %v", service)
	if err != nil {
		logging.Logger.Printf(" 【err】 %v", err)
		response.Fail(c, "注册失败")
	} else {
		response.Success(c, service)
	}

}

// 心跳机制
func (s *Server) heartBeat(c *gin.Context) {

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
func (registry *Registry) heartbeat(serviceName, host string, port int) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	serviceKey := serviceName + host + string(rune(port))
	if service, exists := registry.services[serviceKey]; exists {
		service.LastHeartBeatTime = time.Now().String()
	}
}
