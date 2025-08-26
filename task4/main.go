package main

import (
	"fmt"
	"task4/internal/config"
	"task4/internal/database"
	"task4/internal/model"
	"task4/internal/router"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.SetLevel(logrus.InfoLevel)
	if err := config.Load(); err != nil {
		logrus.WithError(err).Fatal("配置加载失败")
	}
	if err := database.InitMySQL(); err != nil {
		logrus.WithError(err).Fatal("数据库初始化失败")
	}
	if err := database.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}); err != nil {
		logrus.WithError(err).Fatal("自动创建表失败")
	}
	r := router.SetupRouter()
	port := fmt.Sprintf(":%d", config.Cfg.Server.Port)
	logrus.Infof("服务启动 %s", port)
	if err := r.Run(port); err != nil {
		logrus.WithError(err).Fatal("服务启动失败")
	}
}
