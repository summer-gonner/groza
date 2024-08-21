package utils

import (
	"github.com/google/uuid"
)

// GenerateTraceId 用uuid 作为traceId
func GenerateTraceId() string {
	return uuid.New().String()
}
