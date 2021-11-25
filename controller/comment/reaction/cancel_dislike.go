package reaction

import (
	"errors"
	"net/http"
	"strconv"

	"Hackathon/model"
	dto "Hackathon/model/DTO"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CancelDislike() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, err := strconv.Atoi(c.Param("comment_id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := model.DB.Exec("update comment set dislike_count = dislike_count - 1 where id = ?", ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.Status(http.StatusBadRequest)
				return
			} else {
				panic(err)
			}
		}

		if err := model.DB.Delete(&dto.CommentLike{CommentID: uint(ID), UserID: c.GetString("user_id")}).Error; err != nil {
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
