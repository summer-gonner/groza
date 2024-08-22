package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/groza/server"
)

func (r *Routes) ServerRoutes(engine *gin.Engine) {

	s := server.Server{}
	serverGroup := engine.Group("/service")
	{
		//服务注册
		serverGroup.POST("/register", s.ServiceRegister)
		serverGroup.GET("/getAllServices", s.GetAllServices)
		serverGroup.GET("/getService/:serviceName", s.GetService)
		serverGroup.POST("/heartBeat", s.HeartBeat)
	}

	engine.Run(":5566")
}
