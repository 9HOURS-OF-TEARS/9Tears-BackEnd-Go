package sticker

import (
	"net/http"
	"strconv"

	"Hackathon/model"
	dto "Hackathon/model/DTO"
	"github.com/gin-gonic/gin"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, err := strconv.Atoi(c.Param("comment_id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := model.DB.Create(&dto.CommentSticker{
			UserID:    c.GetString("user_id"),
			CommentID: uint(ID),
		}); err != nil {
			panic(err)
		}

		c.Status(http.StatusCreated)
	}
}