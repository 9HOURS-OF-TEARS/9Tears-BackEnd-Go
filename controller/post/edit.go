package post

import (
	"errors"
	"net/http"

	"Hackathon/model"
	dto "Hackathon/model/DTO"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Edit() gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		req := new(request)
		if err := c.ShouldBind(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		ID := c.Param("post_id")

		if err := model.DB.Where("id = ?", ID).Updates(dto.Post{Title: req.Title, Content: req.Content}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.Status(http.StatusBadRequest)
				return
			} else {
				panic(err)
			}
		}

		c.Status(http.StatusOK)
	}
}
