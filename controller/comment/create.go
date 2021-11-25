package comment

import (
	"net/http"
	"strconv"
	"time"

	"Hackathon/model"
	dto "Hackathon/model/DTO"
	"github.com/gin-gonic/gin"
)

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			Content string `json:"content,omitempty"`
		}
		req := new(request)
		PostID, err := strconv.Atoi(c.Param("post_ID"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if err := c.ShouldBind(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		comment := dto.Comment{
			Content:   req.Content,
			UserID:    c.GetString("user_id"),
			PostID:    uint(PostID),
			CreatedAt: time.Now().Add(time.Hour * 9),
		}

		if err := model.DB.Create(&comment).Error; err != nil {
			panic(err)
		}
		c.Status(http.StatusCreated)
	}
}
