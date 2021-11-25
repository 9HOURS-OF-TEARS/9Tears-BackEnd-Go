package main

import (
	"log"
	"net/http"

	"Hackathon/controller/comment"
	"Hackathon/controller/comment/emoji"
	commentAction "Hackathon/controller/comment/reaction"
	commentSticker "Hackathon/controller/comment/sticker"
	"Hackathon/controller/post"
	"Hackathon/controller/post/reaction"
	"Hackathon/controller/post/sticker"
	"Hackathon/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	model.Connect()

	r := gin.New()
	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.Status(http.StatusInternalServerError)
	}))
	r.Use(gin.Logger())

	postAPI := r.Group("/post")
	{
		postAPI.POST("/", post.Create())
		postAPI.GET("/", post.GetList())
		postAPI.GET("/:post_id", post.Get())
		postAPI.PATCH("/:post_id", post.Edit())
		postAPI.DELETE("/:post_id", post.Delete())

		postReactionAPI := postAPI.Group("/:post_id")
		{
			postReactionAPI.POST("/like", reaction.Like())
			postReactionAPI.DELETE("/like", reaction.CancelLike())
			postReactionAPI.POST("/dislike", reaction.Dislike())
			postReactionAPI.DELETE("/dislike", reaction.CancelDislike())

			postReactionAPI.POST("/sticker", sticker.Create())
			postReactionAPI.DELETE("/sticker", sticker.Delete())
		}

		commentAPI := postAPI.Group("/:post_id/comment")
		{
			commentAPI.POST("/", comment.Create())
			commentAPI.DELETE("/:comment_id", comment.Delete())

			commentReactionAPI := commentAPI.Group("/:comment_id")
			{
				commentReactionAPI.POST("/like", commentAction.Like())
				commentReactionAPI.DELETE("/like", commentAction.CancelLike())
				commentReactionAPI.POST("/dislike", commentAction.Dislike())
				commentReactionAPI.DELETE("/dislike", commentAction.CancelDislike())

				commentReactionAPI.POST("/sticker", commentSticker.Create())
				commentReactionAPI.DELETE("/sticker", commentSticker.Delete())

				commentReactionAPI.POST("/emoji", emoji.Create())
				commentReactionAPI.DELETE("/emoji", emoji.Delete())
			}

		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
