package server

import (
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/groza/registry"

	"github.com/summer-gonner/groza/common/response"
	"github.com/summer-gonner/groza/logging"
)

type Server struct {
}

// 服务注册
func (s *Server) GetService(c *gin.Context) {
	name := c.Param("name")
	response.Success(c, name)
}

func (s *Server) GetAllServices(c *gin.Context) {
	registry := registry.Registry{}
	services := registry.GetAllServices()
	logging.Info("查询到的服务信息为 %s", services)
	response.Success(c, services)
}

func (s *Server) ServiceRegister(c *gin.Context) {

	var service registry.ServiceInstance
	err := c.ShouldBindJSON(&service)
	logging.Info(" 【sevice】", service)
	r := registry.Registry{}
	r.RegisterService(&service)
	if err != nil {
		logging.Errorf(" 【err】", err)
		response.Fail(c, "注册失败")
	} else {
		response.Success(c, service)
	}

}

// HeartBeat 心跳机制
func (s *Server) HeartBeat(c *gin.Context) {

}
