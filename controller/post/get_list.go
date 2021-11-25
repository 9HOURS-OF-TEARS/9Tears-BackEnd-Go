package post

import (
	"errors"
	"net/http"

	"Hackathon/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		type post struct {
			Title        string `json:"title"`
			Nickname     string `json:"nickname"`
			CreatedAt    string `json:"created_at"`
			CommentCount uint   `json:"comment_count"`
			LikeCount    uint   `json:"like_count"`
			DislikeCount uint   `json:"dislike_count"`
		}

		type response struct {
			Post []post `json:"post"`
		}
		res := new(response)

		criterion := c.Query("criterion")
		duration := c.DefaultQuery("duration", "daily")

		if criterion == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		subQuery := model.DB.Select("count(comment.id)").Group("post_id").Table("comment")
		query := model.DB.Table("post").Select("post.id, post.title, post.created_at, u.nickname, post.like_count, "+
			"post.dislike_count, (?) as comment_count", subQuery).Joins("join user u on post.user_id = u.id").
			Group("post.id")

		switch criterion {
		case "latest":
			query = query.Order("created_at")
		case "popularity":
			query = query.Order("like_count")
			switch duration {
			case "daily":
				query = query.Where("created_at >= CURDATE() - interval 1 day")
			case "monthly":
				query = query.Where("created_at >= CURDATE() - interval 1 month")
			case "weekly":
				query = query.Where("created_at >= CURDATE() - interval 1 week")
			case "yearly":
				query = query.Where("created_at >= CURDATE() - interval 1 year")
			default:
				c.Status(http.StatusBadRequest)
			}
		default:
			c.Status(http.StatusBadRequest)
		}

		if err := query.Scan(&res.Post).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				panic(err)
			}
		}

		c.JSON(http.StatusOK, res)
	}
}
