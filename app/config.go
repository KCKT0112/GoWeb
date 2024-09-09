package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
}

var AppConfig *Config

// 初始化配置
func InitConfig() {
	viper.SetConfigName("config")       // 配置文件名（不包括扩展名）
	viper.SetConfigType("yaml")         // 配置文件类型
	viper.AddConfigPath("./app/config") // 配置文件路径

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	AppConfig = &Config{}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
