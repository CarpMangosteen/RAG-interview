package repository

import (
	"github.com/carpmangosteen/rag-interview/server/internal/model"
	"gorm.io/gorm"
)

// KnowledgeBaseRepo 定义了知识库仓库的接口
type KnowledgeBaseRepo interface {
	Create(kb *model.KnowledgeBase) error
}

// knowledgeBaseRepo 是 KnowledgeBaseRepo 的 GORM 实现
type knowledgeBaseRepo struct {
	db *gorm.DB
}

// NewKnowledgeBaseRepo 创建一个新的知识库仓库实例
func NewKnowledgeBaseRepo(db *gorm.DB) KnowledgeBaseRepo {
	return &knowledgeBaseRepo{db: db}
}

// Create 在数据库中创建一条新的知识库记录
func (r *knowledgeBaseRepo) Create(kb *model.KnowledgeBase) error {
	return r.db.Create(kb).Error
}
