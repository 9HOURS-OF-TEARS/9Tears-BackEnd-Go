package DTO

type PostSticker struct {
	PostID uint   `json:"post_id,omitempty"`
	UserID string `json:"user_id,omitempty"`
}

func (PostSticker) TableName() string {
	return "post_sticker"
}
