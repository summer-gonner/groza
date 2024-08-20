package groza

import (
	"sync"
)

// Registry 用于存储所有注册的服务
type Registry struct {
	services map[string]*ServiceInstance
	mu       sync.Mutex
}

// Default  NewRegistryDefault 创建一个新的注册中心实例
func Default(vf ...VariableFunc) *Registry {
	debugPrintWARNINGDefault()

}

func Create(vf ...VariableFunc) *Registry {
	debugPrintWARNINGCREATE()
}

type VariableFunc func(*Registry)

func debugPrintWARNINGCREATE() {
	debugPrint("打印debug级别的日志")
}
