package server

import (
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/groza/registry"

	"github.com/summer-gonner/groza/common/response"
	"github.com/summer-gonner/groza/logging"
)

type Server struct {
	registry *registry.Registry
}

// NewServer 构造函数，用于初始化 Server 实例
func NewServer() *Server {
	return &Server{
		registry: registry.NewRegistry(), // 初始化 registry
	}
}

// 服务注册
func (s *Server) GetService(c *gin.Context) {
	name := c.Param("name")
	response.Success(c, name)
}

func (s *Server) GetAllServices(c *gin.Context) {
	services := s.registry.GetAllServices()
	logging.Info("查询到的服务信息为 %s", services)
	response.Success(c, services)
}

func (s *Server) ServiceRegister(c *gin.Context) {

	var serviceInfoRequest registry.ServiceInfoRequest
	err := c.ShouldBindJSON(&serviceInfoRequest)
	var serviceInfoDTO registry.ServiceInfoDTO
	serviceInfoDTO.ServiceName = serviceInfoRequest.ServiceName
	serviceInfoDTO.Host = serviceInfoRequest.Host
	serviceInfoDTO.Port = serviceInfoRequest.Port
	logging.Info(" 【ServiceInfoRequest】", serviceInfoRequest)
	re := s.registry.RegisterService(&serviceInfoDTO)
	if err != nil {
		logging.Errorf(" 【注册失败】", err)
		response.Fail(c, "注册失败")
	} else {
		response.Success(c, re)
	}

}

// HeartBeat 心跳机制
func (s *Server) HeartBeat(c *gin.Context) {

}
