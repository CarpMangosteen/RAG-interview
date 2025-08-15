package router

import (
	"github.com/carpmangosteen/rag-interview/server/internal/handler"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有的路由
func RegisterRoutes(r *gin.Engine, kbHandler *handler.KnowledgeBaseHandler) {
	// 健康检查端点
	r.GET("/ping", handler.Ping)

	// v1 版本的 API 路由组
	apiV1 := r.Group("/api/v1")
	{
		// 知识库相关的路由
		kbRoutes := apiV1.Group("/kb")
		{
			kbRoutes.POST("", kbHandler.CreateKnowledgeBase)
			// 未来可以添加其他路由，如：
			// kbRoutes.GET("", kbHandler.ListKnowledgeBases)
			// kbRoutes.GET("/:id", kbHandler.GetKnowledgeBase)
		}
	}
}
