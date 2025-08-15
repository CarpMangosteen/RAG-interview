package cmd

import (
	"fmt"
	"os"

	"github.com/carpmangosteen/rag-interview/server/conf"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "rag-server",
	Short: "A RAG application server",
	Long:  `A RAG application server built with Cobra, GORM, and Gin.`,
}

// Execute 入口位置
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Execute failed:%v", err)
		os.Exit(1)
	}
}
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./config.yaml)")
}

func initConfig() {
	conf.InitConfig(cfgFile)
}
