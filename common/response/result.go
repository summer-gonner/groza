package response

// Response 统一的返回体结构
type Response struct {
	Code      int         `json:"code"`    // 状态码
	Message   string      `json:"message"` // 提示信息
	Data      interface{} `json:"data"`    // 返回的数据
	Success   bool        `json:"success"`
	Fail      bool        `json:"fail"`
	ErrorMsg  string      `json:"errorMsg"`
	RequestId string      `json:"requestId"`
}

// 成功时使用的状态码和消息
const (
	OK    = 1
	ERROR = 0
)

func StatusText(code int) string {
	switch code {
	case OK:
		return "操作成功"
	case ERROR:
		return "操作失败"
	default:
		return ""
	}
}
