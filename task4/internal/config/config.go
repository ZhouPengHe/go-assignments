package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire int64  `mapstructure:"expire"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Charset  string `mapstructure:"charset"`
}

type Config struct {
	Server   ServerConfig `mapstructure:"server"`
	JWT      JWTConfig    `mapstructure:"jwt"`
	Database MySQLConfig  `mapstructure:"database"`
}

var Cfg *Config

func Load() error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return err
	}
	Cfg = &c
	logrus.WithFields(logrus.Fields{"port": c.Server.Port, "db": c.Database.Name}).Info("配置加载成功")
	return nil
}
