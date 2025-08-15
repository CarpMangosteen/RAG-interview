package cmd

import (
	"log"

	"github.com/carpmangosteen/rag-interview/server/conf"
	"github.com/carpmangosteen/rag-interview/server/internal/db"
	"github.com/carpmangosteen/rag-interview/server/internal/es"
	"github.com/carpmangosteen/rag-interview/server/internal/handler"
	"github.com/carpmangosteen/rag-interview/server/internal/repository"
	"github.com/carpmangosteen/rag-interview/server/internal/server"
	"github.com/carpmangosteen/rag-interview/server/internal/service"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the HTTP server",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	cfg := conf.GetCfg()

	if err := db.InitDB(&cfg.Mysql); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// --- 手动依赖注入 ---
	// 获取数据库实例
	gormDB := db.GetDB()

	if err := es.InitES(&cfg.Es); err != nil {
		log.Fatalf("Failed to initialize elasticsearch: %v", err)
	}

	// 初始化 Repository 层
	kbRepo := repository.NewKnowledgeBaseRepo(gormDB)

	// 初始化 Service 层
	kbSvc := service.NewKnowledgeBaseSvc(kbRepo)

	// 初始化 Handler 层
	kbHandler := handler.NewKnowledgeBaseHandler(kbSvc)

	// 初始化 HTTP 服务，并将 Handler 注入
	httpServer := server.NewHTTPServer(kbHandler)
	// --- 注入完成 ---

	log.Printf("Server starting on %s", cfg.Server.Address)
	if err := httpServer.Run(cfg.Server.Address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
