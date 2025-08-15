package service

import (
	"github.com/carpmangosteen/rag-interview/server/internal/model"
	"github.com/carpmangosteen/rag-interview/server/internal/repository"
	"github.com/carpmangosteen/rag-interview/server/internal/service/dto"
)

// KnowledgeBaseSvc 定义了知识库服务的接口
type KnowledgeBaseSvc interface {
	CreateKnowledgeBase(req *dto.CreateKnowledgeBaseReq) (*model.KnowledgeBase, error)
}

// knowledgeBaseSvc 是 KnowledgeBaseSvc 的实现
type knowledgeBaseSvc struct {
	kbRepo repository.KnowledgeBaseRepo
}

// NewKnowledgeBaseSvc 创建一个新的知识库服务实例
func NewKnowledgeBaseSvc(kbRepo repository.KnowledgeBaseRepo) KnowledgeBaseSvc {
	return &knowledgeBaseSvc{kbRepo: kbRepo}
}

// CreateKnowledgeBase 处理创建知识库的业务逻辑
func (s *knowledgeBaseSvc) CreateKnowledgeBase(req *dto.CreateKnowledgeBaseReq) (*model.KnowledgeBase, error) {
	kb := &model.KnowledgeBase{
		Name: req.Name,
		Desc: req.Desc,
	}

	err := s.kbRepo.Create(kb)
	if err != nil {
		// 这里未来可以处理更复杂的错误，比如名字重复等
		return nil, err
	}

	return kb, nil
}
