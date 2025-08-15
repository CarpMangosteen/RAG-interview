package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *Conf

func InitConfig(cfgFile string) error {
	if cfgFile == "" {
		return fmt.Errorf("not find config file path")
	}

	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("ReadInConfig file failed:%v", err))
	}

	cfg = new(Conf)
	err := viper.Unmarshal(cfg)
	if err != nil {
		panic(fmt.Errorf("unmarshal file failed:%v", err))
	}

	return nil
}

func GetCfg() *Conf {
	if cfg == nil {
		panic("Fatal error config is not init")
	}

	return cfg
}

// Config 结构体与 config.yaml 文件完全对应
type Conf struct {
	Server ServerConfig `mapstructure:"server"`
	Logger LoggerConfig `mapstructure:"logger"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Es     EsConfig     `mapstructure:"es"`
	Model  ModelConfig  `mapstructure:"model"`
}

type ServerConfig struct {
	Address string `mapstructure:"address"`
}

type LoggerConfig struct {
	Level  string `mapstructure:"level"`
	Stdout bool   `mapstructure:"stdout"`
}

type MysqlConfig struct {
	Link string `mapstructure:"link"`
}

type EsConfig struct {
	Endpoints []string `mapstructure:"endpoints"`
	Username  string   `mapstructure:"username"`
	Password  string   `mapstructure:"password"`
}

type ModelConfig struct {
	Embedding ModelInfo `mapstructure:"embedding"`
	Llm       ModelInfo `mapstructure:"llm"`
}

type ModelInfo struct {
	Provider string `mapstructure:"provider"`
	Name     string `mapstructure:"name"`
	Endpoint string `mapstructure:"endpoint"`
	ApiKey   string `mapstructure:"api_key"`
}
