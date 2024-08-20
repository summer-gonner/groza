package groza_server

import (
	"sync"
)

// 其他服务请求注册的请求体
type ServiceInstance struct {
	ServiceName  string `json:"serviceName"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	RegisterTime string `json:"registerTime"`
}

// Registry 用于存储所有注册的服务
type Registry struct {
	services map[string]*ServiceInstance
	Port     string
	mu       sync.Mutex
}
