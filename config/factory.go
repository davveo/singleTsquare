package config

import "github.com/jinzhu/configor"

var DefaultConfig = &Config{
	Databse: DataBaseConfig{
		Type:         "mysql",
		Host:         "localhost",
		Port:         3306,
		User:         "root",
		Password:     "123123",
		DatabaseName: "tsquare",
		MaxIdleConns: 10,
		MaxOpenConns: 10,
	},
	Session:       SessionConfig{},
	IsDevelopment: true,
}

func NewDefaultConfig() *Config {
	return DefaultConfig
}

func NewConfig(configFile string) *Config {
	if configFile != "" {
		config := &Config{}
		configor.Load(config, configFile)
		return config
	}
	return NewDefaultConfig()
}
