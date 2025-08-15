package dto

// CreateKnowledgeBaseReq 定义了创建知识库的请求体
type CreateKnowledgeBaseReq struct {
	Name string `json:"name" binding:"required"`
	Desc string `json:"desc"`
}
