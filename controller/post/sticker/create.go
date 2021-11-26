package sticker

import (
	"errors"
	"net/http"
	"strconv"

	"Hackathon/model"
	dto "Hackathon/model/DTO"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, err := strconv.Atoi(c.Param("post_id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := model.DB.Exec("update post set sticker_count = sticker_count + 1 where id = ?", ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.Status(http.StatusBadRequest)
				return
			} else {
				panic(err)
			}
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
