package post

import (
	"errors"
	"net/http"

	"Hackathon/model"
	dto "Hackathon/model/DTO"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		type response struct {
			Post *dto.Post `json:"post"`
		}
		res := new(response)
		res.Post = new(dto.Post)

		ID := c.Param("post_id")

		if err := model.DB.Exec("update post set view_count = view_count + 1 where id = ?", ID).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				panic(err)
			}
		}

		if err := model.DB.Find(&res.Post, "id = ?", ID).Limit(30).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				panic(err)
			}
		}

		if err := model.DB.Find(&res.Post.PostLike, "post_id = ?", ID).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				panic(err)
			}
		}

		if err := model.DB.Find(&res.Post.PostSticker, "post_id = ?", ID).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				panic(err)
			}
		}

		if err := model.DB.Find(&res.Post.Comment, "post_id = ?", ID).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				panic(err)
			}
		}

		for i := range res.Post.Comment {
			if err := model.DB.Find(&res.Post.Comment[i].CommentEmoji, "post_id = ?", ID).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					panic(err)
				}
			}
			if err := model.DB.Find(&res.Post.Comment[i].CommentSticker, "post_id = ?", ID).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					panic(err)
				}
			}
			if err := model.DB.Find(&res.Post.Comment[i].CommentLike, "post_id = ?", ID).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					panic(err)
				}
			}
		}
		c.JSON(http.StatusOK, res)

	}
}
