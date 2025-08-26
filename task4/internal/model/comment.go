package model

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;comment:主键ID" json:"id"`
	Content   string    `gorm:"type:text;not null;comment:评论内容" json:"content"`
	UserID    uint      `gorm:"not null;index;comment:评论用户ID" json:"user_id"`
	PostID    uint      `gorm:"not null;index;comment:文章ID" json:"post_id"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
}
