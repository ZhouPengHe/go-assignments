package handler

import (
	"task4/internal/database"
	"task4/internal/model"
	"task4/internal/utils"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}

// 创建评论（鉴权）
func CreateComment(c *gin.Context) {
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "请求参数格式错误")
		return
	}
	comment := model.Comment{Content: req.Content, PostID: req.PostID, UserID: c.GetUint("user_id")}
	if err := database.DB.Create(&comment).Error; err != nil {
		utils.Error(c, "创建评论失败")
		return
	}
	utils.Success(c, comment)
}

// 根据文章Id获取评论列表（公开）
func ListComments(c *gin.Context) {
	postID := c.Param("post_id")
	var comments []model.Comment
	if err := database.DB.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		utils.Error(c, "查询失败")
		return
	}
	utils.Success(c, comments)
}

// 删除评论（仅创建人）
func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	var mc model.Comment
	if err := database.DB.First(&mc, id).Error; err != nil {
		utils.Error(c, "评论不存在")
		return
	}
	if mc.UserID != c.GetUint("user_id") {
		utils.Error(c, "无权限删除此评论")
		return
	}
	if err := database.DB.Delete(&mc).Error; err != nil {
		utils.Error(c, "删除失败")
		return
	}
	utils.Success(c, gin.H{"deleted": id})
}
