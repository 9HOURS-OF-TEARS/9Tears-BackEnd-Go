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

func Dislike() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, err := strconv.Atoi(c.Param("comment_id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := model.DB.Exec("update comment set dislike_count = dislike_count + 1 where id = ?", ID).Error; err != nil {
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
			IsLike:    false,
		}); err != nil {
			panic(err)
		}

		var count struct {
			DislikeCount uint `json:"dislike_count,omitempty"`
			LikeCount    uint `json:"like_count,omitempty"`
		}
		if err := model.DB.Select("dislike_count, like_count").Find(&count, "id = ?", ID).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				panic(err)
			}
		}

		if count.DislikeCount/(count.DislikeCount+count.LikeCount)*100 > 60 {
			if err := model.DB.Delete(&dto.Comment{ID: uint(ID)}).Error; err != nil {
				panic(err)
			}
		}
		c.Status(http.StatusCreated)
	}
}
