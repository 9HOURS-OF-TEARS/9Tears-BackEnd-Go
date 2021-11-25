package user

import (
	"net/http"
	"time"

	"Hackathon/model"
	dto "Hackathon/model/DTO"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Clime struct {
	UID string
	Exp *jwt.Time
}

func (c Clime) Valid(helper *jwt.ValidationHelper) error {
	return helper.ValidateExpiresAt(c.Exp)
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			ID       string `json:"id,omitempty"`
			Password string `json:"password,omitempty"`
		}

		req := new(request)
		err := c.ShouldBind(&req)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		var user dto.User
		var count int64
		model.DB.Find(&user, "id = ?", req.ID).Count(&count)

		if count == 0 {
			c.Status(http.StatusBadRequest)
			return
		}

		compare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if compare != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		clime := Clime{
			UID: user.ID,
			Exp: jwt.At(time.Now().Add((time.Hour * 24) * 365 * 10)),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &clime)
		accessToken, err := token.SignedString([]byte("Hackathon2021"))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token": accessToken,
		})

	}
}
