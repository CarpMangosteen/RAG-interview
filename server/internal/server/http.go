package server

import (
	"github.com/carpmangosteen/rag-interview/server/internal/handler"
	"github.com/carpmangosteen/rag-interview/server/internal/router"
	"github.com/gin-gonic/gin"
)

// NewHTTPServer 创建并配置一个新的 Gin HTTP 服务实例
func NewHTTPServer(kbHandler *handler.KnowledgeBaseHandler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// 注册所有路由，并将 handler 传递进去
	router.RegisterRoutes(r, kbHandler)

	return r
}
