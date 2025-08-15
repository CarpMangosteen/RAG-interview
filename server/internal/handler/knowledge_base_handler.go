package handler

import (
	"net/http"

	"github.com/carpmangosteen/rag-interview/server/internal/service"
	"github.com/carpmangosteen/rag-interview/server/internal/service/dto"
	"github.com/gin-gonic/gin"
)

// KnowledgeBaseHandler 封装了知识库相关的 HTTP 处理逻辑
type KnowledgeBaseHandler struct {
	Svc service.KnowledgeBaseSvc
}

// NewKnowledgeBaseHandler 创建一个新的知识库处理器
func NewKnowledgeBaseHandler(svc service.KnowledgeBaseSvc) *KnowledgeBaseHandler {
	return &KnowledgeBaseHandler{Svc: svc}
}

// CreateKnowledgeBase 处理创建知识库的 API 请求
func (h *KnowledgeBaseHandler) CreateKnowledgeBase(c *gin.Context) {
	var req dto.CreateKnowledgeBaseReq
	// 绑定并验证请求体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用 service 层处理业务逻辑
	kb, err := h.Svc.CreateKnowledgeBase(&req)
	if err != nil {
		// 在实际应用中，这里会根据错误类型返回不同的 HTTP 状态码
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusCreated, kb)
}
