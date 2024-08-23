package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/groza/server"
)

func (r *Routes) ServerRoutes(engine *gin.Engine) {

	srv := server.NewServer()

	serverGroup := engine.Group("/service")
	{
		//服务注册
		serverGroup.POST("/register", srv.ServiceRegister)
		serverGroup.GET("/getAllServices", srv.GetAllServices)
		serverGroup.GET("/getService/:serviceName", srv.GetService)
		serverGroup.POST("/heartBeat", srv.HeartBeat)
	}

	engine.Run(":5566")
}
