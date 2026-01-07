package main

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			SocialNetworkSysDsgn
//	@version		1.0
//	@description	API for getting, creating, updating  and deleting posts and other supported stuff

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		posts := v1.Group("/posts")
		{
			posts.GET("/:id", GetPost)
			posts.GET("/", GetPosts)
			posts.POST("/", CreatePost)
			posts.PATCH("/:id", UpdatePost)
			posts.DELETE("/:id", DeletePost)

			posts.GET("/:id/comments", GetPostComments)
			posts.GET("/:id/rating", GetPostRating)
		}

		comments := v1.Group("/comments")
		{
			comments.POST("/", CreateComment)
			comments.PATCH("/:id", UpdateComment)
			comments.DELETE("/:id", DeleteComment)
		}

		search := v1.Group("/search")
		{
			search.GET("/", SearchByLocation)
		}

		locations := v1.Group("/locations")
		{
			locations.GET("/", GetAllLocations)
			locations.GET("/:id/posts", GetPostsByLocation)
		}

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
