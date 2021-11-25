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
		ID, err := strconv.Atoi(c.Param("post_id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := model.DB.Create(&dto.PostSticker{
			UserID: c.GetString("user_id"),
			PostID: uint(ID),
		}); err != nil {
			panic(err)
		}

		c.Status(http.StatusCreated)
	}
}
