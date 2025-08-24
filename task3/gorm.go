package task3

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
type User struct {
	Id        uint   `gorm:"primarykey"`
	Name      string `gorm:"type:varchar(60);not null"`
	Posts     []Post
	PostCount uint
}

type Post struct {
	Id            uint   `gorm:"primarykey"`
	Title         string `gorm:"type:varchar(1000);not null"`
	Content       string `gorm:"type:longtext;not null"`
	UserID        uint   `gorm:"not null"`
	Comments      []Comment
	CommentCount  int    `gorm:"-"`
	CommentStatus string `gorm:"type:varchar(20)"`
}

type Comment struct {
	Id      uint   `gorm:"primarykey"`
	Content string `gorm:"type:longtext;not null"`
	PostID  uint   `gorm:"not null"`
}

func InitGormDB(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		return err
	}
	// 创建用户数据
	users := []User{
		{Name: "张三"},
		{Name: "李四"},
		{Name: "王五"},
	}
	for _, u := range users {
		var posts []Post
		if u.Name == "张三" || u.Name == "王五" {
			posts = append(posts, Post{
				Title:   fmt.Sprintf("Go入门教程 - 主讲人：%s", u.Name),
				Content: "Go语言基础教程内容",
				Comments: []Comment{
					{Content: fmt.Sprintf("非常好，受益匪浅 - 被评价人：%s", u.Name)},
					{Content: fmt.Sprintf("讲解清晰，易于理解 - 被评价人：%s", u.Name)},
				},
			})
		}
		if u.Name == "李四" || u.Name == "王五" {
			posts = append(posts, Post{
				Title:   fmt.Sprintf("Go并发编程 - 主讲人：%s", u.Name),
				Content: "Go并发编程深入分析",
				Comments: []Comment{
					{Content: fmt.Sprintf("期待更深入的内容 - 被评价人：%s", u.Name)},
					{Content: fmt.Sprintf("并发部分非常有挑战 - 被评价人：%s", u.Name)},
				},
			})
		}
		u.Posts = posts
		if err := db.Create(&u).Error; err != nil {
			return err
		}
	}
	return nil
}

// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func QueryUserInfo(db *gorm.DB, userName string) (User, error) {
	user := User{}
	if userName == "" {
		return user, errors.New("用户不能为空")
	}
	err := db.Preload("Posts").Preload("Posts.Comments").First(&user, "name = ?", userName).Error
	return user, err
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func QueryLargestCommentsNum(db *gorm.DB) (Post, error) {
	var post Post
	var result struct {
		Title        string
		CommentCount int
	}
	err := db.Model(&Post{}).
		Select("posts.title, count(comments.post_id) as CommentCount").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.id").
		Order("CommentCount desc").
		Limit(1).
		Scan(&result).Error
	post.Title = result.Title
	post.CommentCount = result.CommentCount
	return post, err
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (post *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).
		Where("id = ?", post.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).
		Error
}

func DeleteComment(db *gorm.DB, commentId uint) error {
	var comment Comment
	if err := db.Find(&comment, commentId).Error; err != nil {
		return err
	}
	return db.Where("id = ?", commentId).Delete(&comment).Error
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (commment *Comment) AfterDelete(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", commment.PostID).Count(&count).Error; err != nil {
		return err
	}
	var commentStatus string
	if count == 0 {
		commentStatus = "无评论"
	} else {
		commentStatus = fmt.Sprintf("有%d条评论", count)
	}
	return tx.Model(&Post{}).
		Where("id = ?", commment.PostID).
		Update("comment_status", commentStatus).
		Error
}
