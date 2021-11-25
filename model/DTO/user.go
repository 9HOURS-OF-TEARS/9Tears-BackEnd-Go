package DTO

import "database/sql"

type User struct {
	ID       string         `gorm:"primaryKey" json:"id,omitempty"`
	Password string         `gorm:"not null" json:"password,omitempty"`
	Nickname sql.NullString `json:"nickname"`

	Post        []Post        `gorm:"foreignKey: UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"post,omitempty"`
	PostLike    []PostLike    `gorm:"foreignKey: UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"post_like,omitempty"`
	PostSticker []PostSticker `gorm:"foreignKey: UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"post_sticker,omitempty"`

	Comment        []Comment        `gorm:"foreignKey: UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comment,omitempty"`
	CommentLike    []CommentLike    `gorm:"foreignKey: UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comment_like,omitempty"`
	CommentSticker []CommentSticker `gorm:"foreignKey: UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comment_sticker,omitempty"`
	CommentEmoji   []CommentEmoji   `gorm:"foreignKey: UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comment_emoji,omitempty"`
}

func (User) TableName() string {
	return "user"
}
