package groza

// ServiceInstance 其他服务请求注册的请求体信息
type ServiceInstance struct {
	ServiceName       string `json:"serviceName"`
	Host              string `json:"host"`
	Port              int    `json:"port"`
	RegisterTime      string `json:"registerTime"`
	LastHeartBeatTime string `json:"LastHeartBeatTime"`
}
