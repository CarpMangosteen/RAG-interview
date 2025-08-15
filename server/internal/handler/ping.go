package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 是一个简单的健康检查处理函数
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
