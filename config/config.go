package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
)

type MongoDB struct {
	Host     string `json:"host"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Config struct {
	MongoDB        MongoDB `json:"mongodb"`
	InfuraUrl      string  `json:"infura_url"`
	ServerPort     string  `json:"server_port"`
	WorkerQuantity int     `json:"worker_quantity"`
}

func GetEnv(key string, defaultValue string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		return defaultValue
	}
	return val
}

func Load() *Config {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	configFileName := "config"
	viper.SetConfigName(configFileName)
	err := viper.ReadInConfig() // Find and read the configs file
	if err != nil {             // Handle errors reading the configs file
		panic(fmt.Errorf("fatal error configs file: %w", err))
	}

	var cfg *Config
	err = viper.Unmarshal(&cfg, func(c *mapstructure.DecoderConfig) {
		c.TagName = "json"
	})
	if err != nil {
		panic(fmt.Errorf("unmashal configs fail: %w", err))
	}
	env := os.Getenv("APP_ENV")
	if env != "local" {
		cfg.InfuraUrl = GetEnv("INFURA_URL", cfg.InfuraUrl)
		cfg.MongoDB.Host = GetEnv("DB_HOST", cfg.MongoDB.Host)
	}
	return cfg
}
