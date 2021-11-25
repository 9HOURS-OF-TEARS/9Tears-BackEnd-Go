package post

import (
	"net/http"
	"time"

	"Hackathon/model"
	dto "Hackathon/model/DTO"
	"github.com/gin-gonic/gin"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			Title       string `json:"title" binding:"required"`
			Content     string `json:"content"`
			IsAnonymous bool   `json:"is_anonymous"`
		}
		req := new(request)

		if err := c.ShouldBind(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		post := dto.Post{
			Title:       req.Title,
			Content:     req.Content,
			UserID:      c.GetString("user_id"),
			IsAnonymous: req.IsAnonymous,
			CreatedAt:   time.Now().Add(time.Hour * 9),
		}

		if err := model.DB.Create(&post).Error; err != nil {
			panic(err)
		}

		c.Status(http.StatusCreated)
	}
}
