package config

import (
	"sync"

	"github.com/spf13/viper"
	"github.com/versegeek/toolkit/pkg/common"
)

const (
	ServerPortKey = "SERVER_PORT"
)

type Config struct {
	ServerPort string
}

func NewConfig() *Config {
	v := viper.New()
	v.SetDefault(ServerPortKey, "8080")

	common.LoadFromFile(v)

	return &Config{
		ServerPort: v.GetString(ServerPortKey),
	}
}

var (
	conf *Config
	once sync.Once
)

func GetConfig() *Config {
	if conf == nil {
		once.Do(func() {
			conf = NewConfig()
		})
	}

	return conf
}
