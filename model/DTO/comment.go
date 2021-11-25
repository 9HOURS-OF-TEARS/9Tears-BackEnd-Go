package DTO

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey; autoIncrement" json:"id,omitempty"`
	Content   string    `json:"content,omitempty"`
	UserID    string    `json:"user_id,omitempty"`
	PostID    uint      `json:"post_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`

	CommentLike    []CommentLike    `gorm:"foreignKey: CommentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comment_like"`
	CommentSticker []CommentSticker `gorm:"foreignKey: CommentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comment_sticker"`
	CommentEmoji   []CommentEmoji   `gorm:"foreignKey: CommentID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comment_emoji"`
}

func (Comment) TableName() string {
	return "comment"
}
