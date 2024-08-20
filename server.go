package groza_server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 启动服务
func (registry *Registry) StartServer(port string) {
	r := gin.Default()
	// 服务注册接口
	r.POST("/register", func(c *gin.Context) {
		var service ServiceInstance
		if err := c.ShouldBindJSON(&service); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Print("注册进来的服务信息%s", service)
		registry.registerService(&service)
		c.JSON(http.StatusOK, gin.H{"status": "registered"})
	})

	// 查询所有服务
	r.GET("/services", func(c *gin.Context) {
		services := registry.getAllServices()
		c.JSON(http.StatusOK, services)
	})

	// 查询特定服务
	r.GET("/service/:name", func(c *gin.Context) {
		name := c.Param("name")
		service := registry.getService(name)
		if service == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "服务没找到"})
			return
		}
		c.JSON(http.StatusOK, service)
	})
	err := r.Run(port)
	if err != nil {
		return
	}
}

// RegisterService 注册一个服务
func (registry *Registry) registerService(service *ServiceInstance) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	serviceKey := service.ServiceName + service.Host + string(service.Port)
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
func (r *Registry) heartbeat(serviceName, host string, port int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	serviceKey := serviceName + host + string(port)
	if service, exists := r.services[serviceKey]; exists {
		service.LastHeartBeatTime = time.Now().String()
	}
}
