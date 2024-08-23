package registry

// ServiceInstance 其他服务请求注册的请求体信息
type ServiceInfoResponse struct {
	ServiceName       string `json:"serviceName"`
	Host              string `json:"host"`
	Port              int    `json:"port"`
	RegisterTime      string `json:"registerTime"`
	LastHeartBeatTime string `json:"lastHeartBeatTime"`
}
type ServiceInfoRequest struct {
	ServiceName string `json:"serviceName"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
}
type ServiceInfoDTO struct {
	ServiceName       string `json:"serviceName"`
	Host              string `json:"host"`
	Port              int    `json:"port"`
	RegisterTime      string `json:"registerTime"`
	IsOnline          bool   `json:"isOnline"`
	LastHeartBeatTime string `json:"lastHeartBeatTime"`
}
type ServiceInfoVoList []*ServiceInfoDTO
