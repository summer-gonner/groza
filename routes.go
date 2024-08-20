package groza_server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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
		registry.RegisterService(&service)
		c.JSON(http.StatusOK, gin.H{"status": "registered"})
	})

	// 查询所有服务
	r.GET("/services", func(c *gin.Context) {
		services := registry.GetAllServices()
		c.JSON(http.StatusOK, services)
	})

	// 查询特定服务
	r.GET("/service/:name", func(c *gin.Context) {
		name := c.Param("name")
		service := registry.GetService(name)
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
