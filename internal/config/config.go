package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env     string        `yaml:"env"`
	GRPC    gRPCConfig    `yaml:"grpc_server"`
	Storage StorageConfig `yaml:"database"`
}
type gRPCConfig struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
}

type StorageConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func MustLoad() *Config {
	var config Config

	configPath := os.Getenv("config_path")

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		panic(err)
	}

	return &config
}
