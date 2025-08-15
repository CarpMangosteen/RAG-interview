package es

import (
	"log"

	"github.com/carpmangosteen/rag-interview/server/conf"
	"github.com/elastic/go-elasticsearch/v8"
)

var esClient *elasticsearch.Client

// InitES 初始化 Elasticsearch 客户端
func InitES(cfg *conf.EsConfig) error {
	var err error
	esCfg := elasticsearch.Config{
		Addresses: cfg.Endpoints,
		Username:  cfg.Username,
		Password:  cfg.Password,
	}
	esClient, err = elasticsearch.NewClient(esCfg)
	if err != nil {
		return err
	}

	// 检查与 ES 的连接
	res, err := esClient.Info()
	if err != nil {
		return err
	}
	defer res.Body.Close()
	log.Println("Elasticsearch client initialized successfully.")
	return nil
}

// GetES 返回一个 Elasticsearch 客户端实例
func GetES() *elasticsearch.Client {
	if esClient == nil {
		panic("elasticsearch client not initialized")
	}
	return esClient
}
