package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server    ServerConfig    `yaml:"server"`
	Database  DatabaseConfig  `yaml:"database"`
	Redis     RedisConfig     `yaml:"redis"`
	JWT       JWTConfig       `yaml:"jwt"`
	AliyunSMS AliyunSMSConfig `yaml:"aliyun_sms"`
	SIWE      SIWEConfig      `yaml:"siwe"`
}

type ServerConfig struct {
	Addr  string `yaml:"addr"`
	Debug bool   `yaml:"debug"`
}

type DatabaseConfig struct {
	DSN string `yaml:"dsn"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type JWTConfig struct {
	Secret   string `yaml:"secret"`
	ExpHours int    `yaml:"exp_hours"`
}

type AliyunSMSConfig struct {
	AccessKeyID     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	SignName        string `yaml:"sign_name"`
	TemplateCode    string `yaml:"template_code"`
	DevMode         bool   `yaml:"dev_mode"`
}

type SIWEConfig struct {
	ChainID int64 `yaml:"chain_id"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if cfg.Server.Addr == "" {
		cfg.Server.Addr = ":8080"
	}
	if cfg.JWT.ExpHours == 0 {
		cfg.JWT.ExpHours = 168
	}
	if cfg.SIWE.ChainID == 0 {
		cfg.SIWE.ChainID = 1
	}
	if cfg.Redis.Password == "" {
		return nil, fmt.Errorf("redis.password is required")
	}

	return &cfg, nil
}
