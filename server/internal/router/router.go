package router

import (
	"github.com/carpmangosteen/rag-interview/server/internal/handler"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有的路由
func RegisterRoutes(r *gin.Engine) {
	// 一个简单的健康检查端点
	r.GET("/ping", handler.Ping)

	// 为 v1 版本的 API 创建一个路由组
	// apiV1 := r.Group("/api/v1")
	{
		// 未来我们的知识库、聊天等接口都会在这里注册
		// 例如: apiV1.POST("/kb", handler.CreateKnowledgeBase)
	}
}
