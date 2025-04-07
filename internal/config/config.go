package config

import "os"

type Config struct {
	Env  string
	Addr string
	Port string
}

func MustLoad() *Config {
	return &Config{
		Env:  os.Getenv("env"),
		Addr: os.Getenv("addr"),
		Port: os.Getenv("port"),
	}
}
