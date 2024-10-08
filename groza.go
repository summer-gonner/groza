package groza

import (
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/groza/logging"
	"github.com/summer-gonner/groza/middleware"
	"github.com/summer-gonner/groza/registry"
	routes "github.com/summer-gonner/groza/routes"
)

// Default  NewRegistryDefault 创建一个新的注册中心实例
func Default(vf ...VariableFunc) {
	initDefaultGinNew()
	debugPrintWARNINGDefault("哈哈")
}

func Create(vf ...VariableFunc) {
	initDefaultGinNew()
	logging.Logger.Info("groza注册中心启动成功")
	debugPrintWARNINGCREATE()
}

func initDefaultGinNew() {

	engine := gin.New()
	engine.StaticFile("/", "www/index.html")
	engine.Use(middleware.HttpIntercept())
	routers := routes.Routes{}
	routers.ServerRoutes(engine)

}

type VariableFunc func(*registry.Registry)

func debugPrintWARNINGCREATE() {
	debugPrint("打印debug级别的日志")
}

func Use(middleware ...HandlerFunc) {

}

type HandlerFunc func()
