package registry

import (
	"github.com/summer-gonner/groza/logging"
	"github.com/summer-gonner/groza/utils"
	"sync"
	"time"
)

// Registry 用于存储所有注册的服务
type Registry struct {
	services map[string]*ServiceInfoDTO
	mu       sync.Mutex
}

func NewRegistry() *Registry {
	return &Registry{
		services: make(map[string]*ServiceInfoDTO), // 初始化 services 为一个空的 map
	}
}

// StartServer 启动服务
func (registry *Registry) StartServer(port string) {

}

// RegisterService 注册一个服务
func (registry *Registry) RegisterService(serviceInfoDTO *ServiceInfoDTO) *ServiceInfoResponse {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	if registry.services == nil {
		registry.services = make(map[string]*ServiceInfoDTO) // 初始化 map
	}
	format := time.Now().Format("2006-01-02 15:04:05")
	serviceInfoDTO.RegisterTime = format
	serviceInfoDTO.LastHeartBeatTime = format
	//用服务名作为key
	registry.services[serviceInfoDTO.ServiceName] = serviceInfoDTO
	var serviceInfoResponse ServiceInfoResponse
	serviceInfoResponse.LastHeartBeatTime = format
	serviceInfoResponse.RegisterTime = format
	serviceInfoResponse.ServiceName = serviceInfoDTO.ServiceName
	serviceInfoResponse.Host = serviceInfoDTO.Host
	serviceInfoResponse.Port = serviceInfoDTO.Port

	logging.Info("注册的服务所有为 %s", registry.services)
	//for _, ser := range registry.services {
	//	serviceInfoResponse.AllRegisteredService = append(serviceInfoResponse.AllRegisteredService, ser)
	//}
	return &serviceInfoResponse
}

// GetService 查询注册的服务
func (registry *Registry) getService(name, host string, port int) *ServiceInfoResponse {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	//var serviceInfoList ServiceInfoList
	//for _, ser := range registry.services {
	//	serviceInfoList = append(serviceInfoList, ser)
	//}
	return nil
}

// GetAllServices 查询所有注册的服务
func (registry *Registry) GetAllServices() *ServiceInfoVoList {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	var serviceInfoVoList ServiceInfoVoList
	for _, ser := range registry.services {
		//判断最后一次心跳的时间和当前时间  如果超过3分钟 那么视为该服务已经宕机
		compareTimeResult := utils.CompareTime(ser.LastHeartBeatTime, 3)
		logging.Info("比较的时间结果", compareTimeResult)
		ser.IsOnline = compareTimeResult
		serviceInfoVoList = append(serviceInfoVoList, ser)
	}
	return &serviceInfoVoList
}

// Heartbeat 更新服务实例的心跳时间
func (registry *Registry) Heartbeat(serviceName, host string, port int) {
	registry.mu.Lock()
	defer registry.mu.Unlock()

	//serviceKey := serviceName + host + strconv.Itoa(port)
	//if service, exists := registry.services[serviceKey]; exists {
	//	service = time.Now().String()
	//}
}
