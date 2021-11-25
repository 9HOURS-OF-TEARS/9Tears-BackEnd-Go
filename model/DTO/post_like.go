package DTO

type PostLike struct {
	UserID string `json:"user_id,omitempty"`
	PostID uint   `json:"post_id,omitempty"`
	IsLike bool   `json:"is_like,omitempty"`
}

func (PostLike) TableName() string {
	return "post_like"
}
