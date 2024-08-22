package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/groza"
)

func (r *Routes) ServerRoutes(engine *gin.Engine) {

	server := groza.Server{}
	serverGroup := engine.Group("/service")
	{
		//服务注册
		serverGroup.POST("/register", server.ServiceRegister)
		serverGroup.GET("/getAllServices", server.GetAllServices)
		serverGroup.GET("/getService/:serviceName", server.GetService)
		serverGroup.POST("/heartBeat", server.HeartBeat)
	}

	engine.Run(":5566")
}
