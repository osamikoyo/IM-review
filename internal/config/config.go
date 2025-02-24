package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct{
	Host string `yaml:"host"`
	Port uint32 `yaml:"port"`
	Dsn string `yaml:"dsn"`
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil{
		return nil, err
	}

	body, err := io.ReadAll(file)
	if err != nil{
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(body, &cfg)
	return &cfg, err
}