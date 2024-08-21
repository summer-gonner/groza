package groza

import (
	"github.com/gin-gonic/gin"
)

func (r *Routes) ServerRoutes(engine *gin.Engine) {

	server := Server{}
	serverGroup := engine.Group("/service")
	{
		//服务注册
		serverGroup.POST("/register", server.serviceRegister)
		serverGroup.GET("/getAllServices", server.getAllServices)
		serverGroup.GET("/getService/:serviceName", server.getService)
		serverGroup.POST("/heartBeat", server.heartBeat)
	}
	//	// 服务注册接口
	//	engine.POST("/register", func(c *gin.Context) {
	//		var service ServiceInstance
	//		if err := c.ShouldBindJSON(&service); err != nil {
	//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//			return
	//		}
	//		log.Printf("注册进来的服务信息 %s", service)
	//		registry.registerService(&service)
	//		c.JSON(http.StatusOK, gin.H{"status": "registered"})
	//	})
	//
	//	// 查询所有服务
	//	engine.GET("/services", func(c *gin.Context) {
	//		services := registry.getAllServices()
	//		c.JSON(http.StatusOK, services)
	//	})
	//
	//	// 查询特定服务
	//	engine.GET("/service/:name", func(c *gin.Context) {
	//		name := c.Param("name")
	//		service := registry.getService(name)
	//		if service == nil {
	//			c.JSON(http.StatusNotFound, gin.H{"error": "服务没找到"})
	//			return
	//		}
	//		c.JSON(http.StatusOK, service)
	//	})
	//}
	//
	//err := engine.Run(port)
	//if err != nil {
	//	return
	//}
	engine.Run(":5566")
}
