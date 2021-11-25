package comment

import (
	"errors"
	"net/http"

	"Hackathon/model"
	dto "Hackathon/model/DTO"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID := c.Param("comment_id")

		if err := model.DB.Delete(&dto.Comment{}, ID).Error; err != nil {
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
