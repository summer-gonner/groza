package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func response(c *gin.Context, code int, message string, data interface{}, success bool, fail bool, errorMsg string) {
	requestId := c.MustGet("requestId").(string)

	resp := Response{
		Code:      code,
		Message:   message,
		Data:      data,
		Success:   success,
		Fail:      fail,
		ErrorMsg:  errorMsg,
		RequestId: requestId,
	}
	c.JSON(http.StatusOK, resp)
}

// Success 返回成功的响应体
func Success(c *gin.Context, data interface{}) {
	response(c, OK, StatusText(OK), data, true, false, "")
}

// Fail 返回错误的响应体
func Fail(c *gin.Context, errorMsg string) {
	response(c, ERROR, StatusText(ERROR), nil, false, true, errorMsg)
}
