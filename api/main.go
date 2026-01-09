package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/AnnDutova/system_design/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/AnnDutova/system_design/api/rest"
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
		users := v1.Group("/users")
		{
			users.GET("/:id", rest.GetUser)
			users.POST("/", rest.CreateUser)
			users.PATCH("/:id", rest.UpdateUser)
			users.DELETE("/:id", rest.DeleteUser)

			users.PATCH("/:id/subscribe", rest.SubscribeToUser)
			users.PATCH("/:id/unsubscribe", rest.UnsubscribeToUser)
		}

		followers := users.Group("/followers")
		{
			followers.GET("/:id", rest.GetAllUserFollowers)
		}

		posts := v1.Group("/posts")
		{
			posts.GET("/:id", rest.GetPost)
			posts.GET("/", rest.GetPosts)
			posts.POST("/", rest.CreatePost)
			posts.PATCH("/:id", rest.UpdatePost)
			posts.DELETE("/:id", rest.DeletePost)

			posts.GET("/:id/comments", rest.GetPostComments)
			posts.GET("/:id/rating", rest.GetPostRating)
		}

		rating := v1.Group("/rating")
		{
			rating.POST("/", rest.SetRating)
			rating.PATCH("/:id", rest.UpdateRating)
			rating.DELETE("/:id", rest.DeleteRating)
		}

		comments := v1.Group("/comments")
		{
			comments.POST("/", rest.CreateComment)
			comments.PATCH("/:id", rest.UpdateComment)
			comments.DELETE("/:id", rest.DeleteComment)
		}

		locations := v1.Group("/locations")
		{
			locations.POST("/", rest.CreateLocations)
			locations.PATCH("/:id", rest.UpdateLocations)
			locations.DELETE("/:id", rest.DeleteLocations)

			locations.GET("/", rest.GetAllLocations)
			locations.GET("/:id/posts", rest.GetPostsByLocation)
			locations.GET("/filter-fields", rest.SearchByLocation)
		}

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
