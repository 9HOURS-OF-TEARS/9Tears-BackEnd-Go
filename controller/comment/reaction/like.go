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

func Like() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, err := strconv.Atoi(c.Param("comment_id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := model.DB.Exec("update comment set like_count = like_count + 1 where id = ?", ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.Status(http.StatusBadRequest)
				return
			} else {
				panic(err)
			}
		}

		if err := model.DB.Create(&dto.CommentLike{
			UserID:    c.GetString("user_id"),
			CommentID: uint(ID),
			IsLike:    true,
		}); err != nil {
			panic(err)
		}

		c.Status(http.StatusCreated)
	}
}