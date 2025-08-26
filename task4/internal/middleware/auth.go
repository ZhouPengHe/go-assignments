package middleware

import (
	"strings"
	"task4/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 解析 Bearer Token，设置 user_id
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
			utils.Error(c, "未提供有效的Token")
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authorization, "Bearer ")
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			logrus.WithError(err).Error("Token解析失败")
			utils.Error(c, "无效的Token")
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
