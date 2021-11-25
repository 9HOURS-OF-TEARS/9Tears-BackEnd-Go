package emoji

import (
	"net/http"

	"Hackathon/model"
	dto "Hackathon/model/DTO"
	"github.com/gin-gonic/gin"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			CommentID uint   `json:"comment,omitempty"`
			Emoji     string `json:"emoji,omitempty"`
		}
		req := new(request)
		err := c.ShouldBind(&req)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := model.DB.Create(&dto.CommentEmoji{
			UserID:    c.GetString("user_id"),
			CommentID: req.CommentID,
			Emoji:     req.Emoji,
		}); err != nil {
			panic(err)
		}

		c.Status(http.StatusCreated)
	}
}
