package middleware

import (
	"errors"
	"net/http"

	"Hackathon/controller/user"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim := new(user.Clime)
		_, err := jwt.ParseWithClaims(c.GetHeader("Authorization"), claim, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				ErrUnexpectedSigningMethod := errors.New("unexpected signing method")
				return nil, ErrUnexpectedSigningMethod
			}
			return []byte("Hackathon2021"), nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}

		c.Set("user_id", claim.UID)
		c.Next()
	}
}
