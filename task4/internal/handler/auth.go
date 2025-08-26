package handler

import (
	"task4/internal/database"
	"task4/internal/model"
	"task4/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type req struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email;" binding:"omitempty,email"`
}

// 注册
func Register(c *gin.Context) {
	var req req
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "请求参数格式有误")
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := model.User{Username: req.Username, Password: string(hash), Email: req.Email}
	if err := database.DB.Create(&user).Error; err != nil {
		logrus.WithError(err).Error("用户注册失败")
		utils.Error(c, "用户创建失败，用户名或邮箱已存在")
		return
	}
	utils.Success(c, gin.H{"userId": user.ID})
}

// 登录
func Login(c *gin.Context) {
	var req req
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "请求参数格式有误")
		return
	}
	var user model.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.Error(c, "用户不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.Error(c, "用户名或密码不正确")
		return
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		logrus.WithError(err).Error("Token 生成失败")
		utils.Error(c, "系统错误，请稍后重试")
		return
	}
	utils.Success(c, gin.H{"token": token})
}
