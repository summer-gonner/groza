package groza_server

import (
	"sync"
	"time"
)

type ServiceInstance struct {
	ServiceName  string    `json:"serviceName"`
	Host         string    `json:"host"`
	Port         int       `json:"port"`
	RegisterTime time.Time `json:"registerTime"`
}

// Registry 用于存储所有注册的服务
type Registry struct {
	services map[string]*ServiceInstance
	mu       sync.Mutex
}
