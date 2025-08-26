package model

import "time"

type Post struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT;comment:主键ID" json:"id"`
	Title     string    `gorm:"type:varchar(200);not null;comment:标题" json:"title"`
	Content   string    `gorm:"type:text;not null;comment:内容" json:"content"`
	UserID    uint      `gorm:"not null;index;comment:用户ID" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
}
