package DTO

type CommentSticker struct {
	CommentID uint   `json:"comment_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
}

func (CommentSticker) TableName() string {
	return "comment_sticker"
}
