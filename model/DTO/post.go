package DTO

import "time"

type Post struct {
	ID           uint      `gorm:"primaryKey; autoIncrement" json:"id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Content      string    `json:"content,omitempty"`
	UserID       string    `json:"user_id,omitempty"`
	IsAnonymous  bool      `json:"is_anonymous,omitempty"`
	LikeCount    uint      `json:"like_count,omitempty"`
	DislikeCount uint      `json:"dislike_count,omitempty"`
	CreatedAt    time.Time `json:"created_at"`

	Comment []Comment `gorm:"foreignKey: PostID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comment"`

	PostLike    []PostLike    `gorm:"foreignKey: PostID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"post_like"`
	PostSticker []PostSticker `gorm:"foreignKey: PostID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"post_sticker"`
}

func (Post) TableName() string {
	return "post"
}
