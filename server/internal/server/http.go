package server

import (
	"github.com/carpmangosteen/rag-interview/server/internal/router"
	"github.com/gin-gonic/gin"
)

// NewHTTPServer 创建并配置一个新的 Gin HTTP 服务实例
func NewHTTPServer() *gin.Engine {
	// 设置 Gin 的模式，可以根据配置来决定是 debug 还是 release
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// 使用一些通用的中间件
	r.Use(gin.Logger())   // 记录请求日志
	r.Use(gin.Recovery()) // 从 panic 中恢复

	// 注册所有路由
	router.RegisterRoutes(r)

	return r
}
