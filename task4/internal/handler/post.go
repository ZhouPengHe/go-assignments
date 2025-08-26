package handler

import (
	"task4/internal/database"
	"task4/internal/model"
	"task4/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PostReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 创建文章（鉴权）
func CreatePost(c *gin.Context) {
	var req PostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "请求参数格式错误")
		return
	}
	post := model.Post{Title: req.Title, Content: req.Content, UserID: c.GetUint("user_id")}
	if err := database.DB.Create(&post).Error; err != nil {
		logrus.WithError(err).Error("创建文章失败")
		utils.Error(c, "创建文章失败")
		return
	}
	utils.Success(c, post)
}

// 文章详情（公开）
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		utils.Error(c, "文章不存在")
		return
	}
	utils.Success(c, post)
}

// 文章列表（公开）
func ListPosts(c *gin.Context) {
	logrus.Info("ListPosts 入参: 无")
	var posts []model.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		utils.Error(c, "查询失败")
		return
	}
	utils.Success(c, posts)
}

// 更新文章（仅创建人）
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var req PostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "请求参数格式错误")
		return
	}
	var post model.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		utils.Error(c, "文章不存在")
		return
	}
	if post.UserID != c.GetUint("user_id") {
		utils.Error(c, "无权限修改此文章")
		return
	}
	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}
	if err := database.DB.Save(&post).Error; err != nil {
		utils.Error(c, "更新失败")
		return
	}
	utils.Success(c, post)
}

// 删除文章（仅作者）
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		utils.Error(c, "文章不存在")
		return
	}
	if post.UserID != c.GetUint("user_id") {
		utils.Error(c, "无权限删除此文章")
		return
	}
	if err := database.DB.Delete(&post).Error; err != nil {
		utils.Error(c, "删除失败")
		return
	}
	utils.Success(c, gin.H{"deleted": id})
}
