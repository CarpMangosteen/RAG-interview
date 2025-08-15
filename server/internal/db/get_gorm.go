package db

import (
	"github.com/carpmangosteen/rag-interview/server/conf"
	"github.com/carpmangosteen/rag-interview/server/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// InitDB 初始化数据库连接
func InitDB(cfg *conf.MysqlConfig) error {
	var err error
	var gormLogger logger.Interface

	// 可以根据配置决定日志级别
	if gin.Mode() == gin.DebugMode {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default.LogMode(logger.Silent)
	}

	db, err = gorm.Open(mysql.Open(cfg.Link), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return err
	}

	// 自动迁移，GORM 会自动创建表
	err = db.AutoMigrate(&model.KnowledgeBase{}, &model.Document{})
	if err != nil {
		return err
	}

	logrus.Info("Database connection initialized and tables migrated.")
	return nil
}

// GetDB 返回一个数据库连接实例
func GetDB() *gorm.DB {
	if db == nil {
		logrus.Panic("database not initialized")
	}
	return db
}
