package DTO

type CommentLike struct {
	IsLike    bool   `json:"is_like,omitempty"`
	CommentID uint   `json:"comment_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
}

func (CommentLike) TableName() string {
	return "comment_like"
}
