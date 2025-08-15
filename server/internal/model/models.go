package model

import "gorm.io/gorm"

// KnowledgeBase 对应数据库中的 knowledge_bases 表
type KnowledgeBase struct {
	gorm.Model        // 包含了 ID, CreatedAt, UpdatedAt, DeletedAt 字段
	Name       string `gorm:"type:varchar(100);uniqueIndex;not null;comment:知识库名称"`
	Desc       string `gorm:"type:varchar(255);comment:知识库描述"`
	// 如果需要，未来可以添加其他字段，比如 Avatar, OwnerID 等
}

// TableName 自定义表名
func (KnowledgeBase) TableName() string {
	return "knowledge_bases"
}

// Document 对应数据库中的 documents 表
type Document struct {
	gorm.Model
	Name            string `gorm:"type:varchar(255);not null;comment:文档名称"`
	KnowledgeBaseID uint   `gorm:"not null;comment:所属知识库ID"`
	// 可以在这里添加更多元数据，如文件大小、类型等
}

// TableName 自定义表名
func (Document) TableName() string {
	return "documents"
}
