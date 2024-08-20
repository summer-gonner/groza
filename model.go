package groza_server

import (
	"sync"
)

// ServiceInstance 其他服务请求注册的请求体信息
type ServiceInstance struct {
	ServiceName       string `json:"serviceName"`
	Host              string `json:"host"`
	Port              string `json:"port"`
	RegisterTime      string `json:"registerTime"`
	LastHeartBeatTime string `json:"LastHeartBeatTime"`
}

// Registry 用于存储所有注册的服务
type Registry struct {
	services map[string]*ServiceInstance
	mu       sync.Mutex
}
