package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configuration
type Configuration struct {
	Server   ServerConf
	Database DatabaseConf
	Storage  StorageConf
}

// DatabaseConf
type DatabaseConf struct {
	Name   string
	User   string
	Passwd string
}

// StorageConf
type StorageConf struct {
	Bucket    string
	AccessKey string
	SecretKey string
}

// ServerConf
type ServerConf struct {
	Port string
}

func ReadConfiguration() (config Configuration) {
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)

	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return config
}
