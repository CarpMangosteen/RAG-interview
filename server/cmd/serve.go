package cmd

import (
	"log"

	"github.com/carpmangosteen/rag-interview/server/conf"
	"github.com/carpmangosteen/rag-interview/server/internal/db"
	"github.com/carpmangosteen/rag-interview/server/internal/server"
	"github.com/spf13/cobra"
)

// 新增一个 serve 子命令
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. 获取加载好的配置
		cfg := conf.GetCfg()

		// 2. 初始化数据库连接 (新增)
		if err := db.InitDB(&cfg.Mysql); err != nil {
			log.Fatalf("Failed to initialize database: %v", err)
		}

		// 3. 初始化 HTTP 服务
		httpServer := server.NewHTTPServer()

		// 4. 启动服务
		log.Printf("Server starting on %s", cfg.Server.Address)
		if err := httpServer.Run(cfg.Server.Address); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
