package main

import (
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/groza"
	"github.com/summer-gonner/groza/middleware"
)

func main() {
	r := gin.Default()
	middleware.InitLogging()
	trace := middleware.Trace{}
	r.Use(trace.RequestId())
	ri := &groza.RouteInfo{
		Engine: r,
	}
	ri.RouteGroup()
}
