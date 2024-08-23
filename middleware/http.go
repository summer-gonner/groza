package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/groza/logging"
	"github.com/summer-gonner/groza/utils"
	"io/ioutil"
	"time"
)

// HttpIntercept  中间件，用于拦截请求的入参和返回的响应体
func HttpIntercept() gin.HandlerFunc {

	return func(c *gin.Context) {
		traceId := utils.GenerateTraceId()
		logging.Init(traceId)
		// 记录请求入参
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // 读取后重新填充Body
		logging.Info("【请求方式】 %s 【请求URL】%s【请求入参】 %s ", c.Request.Method, c.Request.URL, string(bodyBytes))

		//塞入requestID
		c.Set("requestId", traceId)
		// 创建一个记录器来捕获响应体
		responseBody := &responseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = responseBody

		// 记录响应时间
		startTime := time.Now()
		c.Next() // 继续处理请求

		// 记录响应体
		logging.Info("【响应体】 ", responseBody.body.String())
		logging.Info("【消耗时间】", time.Since(startTime))
	}
}

// responseWriter 是一个包装器，用于捕获响应体
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // 捕获响应体
	return w.ResponseWriter.Write(b)
}
