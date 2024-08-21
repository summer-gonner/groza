package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Trace struct {
	Id string `json:"id"`
}

func (r *Trace) RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("requestId", requestID)
		Logging.WithField("【traceId】", requestID)
		c.Next()
	}
}
