package DTO

type CommentEmoji struct {
	CommentID uint   `json:"comment_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	Emoji     string `json:"emoji,omitempty"`
}

func (CommentEmoji) TableName() string {
	return "comment_emoji"
}
