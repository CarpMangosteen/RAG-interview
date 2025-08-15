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
