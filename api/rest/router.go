package rest

import (
	"github.com/gin-gonic/gin"
)

// Users

// GetUser godoc
//
//	@Summary		Get user by id
//	@Description	Return user with all meta info by specific id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"user ID"
//
//	@Success		200	{object}	User
//
//	@Failure		400	{object}	RestError
//	@Failure		422	{object}	RestError
//	@Failure		500	{object}	RestError
//	@Failure		502	{object}	RestError
//
//	@Router			/users/{id} [get]
func GetUser(c *gin.Context) {
	// some logic
}

// CreateUser godoc
//
//	@Summary		Create new users
//	@Description	Create users with all filled fields
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreateUserModel	true	"CreateUserModel object"
//
//	@Success		201		{object}	User
//
//	@Failure		400		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/users/ [post]
func CreateUser(c *gin.Context) {
	// some logic
}

// UpdateUser godoc
//
//	@Summary		Update user by id
//	@Description	Return updated user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		UpdateUserModel	true	"UpdateUserModel object"
//
//	@Success		202		{object}	User
//
//	@Failure		400		{object}	RestError
//	@Failure		404		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/users/{id} [patch]
func UpdateUser(_ *gin.Context) {
	// some logic
}

// DeleteUser godoc
//
//	@Summary		Delete users by id
//	@Description	Delete users by specific id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"user ID"
//
//	@Success		204
//
//	@Failure		400	{object}	RestError
//	@Failure		404	{object}	RestError
//	@Failure		409	{object}	RestError
//	@Failure		500	{object}	RestError
//	@Failure		502	{object}	RestError
//
//	@Router			/users/{id} [delete]
func DeleteUser(_ *gin.Context) {
	// some logic
}

// SubscribeToUser godoc
//
//	@Summary		Subscribe To User
//	@Description	Return Subscribe To User Model
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		SubscriptionModel	true	"SubscriptionModel object"
//
//	@Success		202		{object}	Subscription
//
//	@Failure		400		{object}	RestError
//	@Failure		404		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/users/{id}/subscribe [patch]
func SubscribeToUser(c *gin.Context) {
	// some logic
}

// UnsubscribeToUser godoc
//
//	@Summary		Unsubscribe To User
//	@Description	Return Unsubscribe To User Model
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		SubscriptionModel	true	"SubscriptionModel object"
//
//	@Success		202		{object}	Subscription
//
//	@Failure		400		{object}	RestError
//	@Failure		404		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/users/{id}/unsubscribe [patch]
func UnsubscribeToUser(c *gin.Context) {
	// some logic
}

// GetAllUserFollowers godoc
//
//	@Summary		Get all user followers
//	@Description	Return all info about user followers
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			limit	path		int	true	"limit for page view"
//	@Param			offset	path		int	true	"offset for page view"
//	@Param			id		path		int	true	"user ID"
//
//	@Success		200		{array}		User
//
//	@Failure		400		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/users/followers/{id} [get]
func GetAllUserFollowers(c *gin.Context) {
	// some logic
}

// Posts

// GetPost godoc
//
//	@Summary		Get post by id
//	@Description	Return post with all meta info by specific id
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"post ID"
//
//	@Success		200	{object}	Post
//
//	@Failure		400	{object}	RestError
//	@Failure		422	{object}	RestError
//	@Failure		500	{object}	RestError
//	@Failure		502	{object}	RestError
//
//	@Router			/posts/{id} [get]
func GetPost(c *gin.Context) {
	// some logic
}

// GetPosts godoc
//
//	@Summary		Get posts
//	@Description	Return all posts by limit (default 25) and offset
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			limit	path		int	true	"limit for page view"
//	@Param			offset	path		int	true	"offset for page view"
//
//	@Success		200		{array}		Post
//
//	@Failure		400		{object}	RestError
//	@Failure		404		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/posts/ [get]
func GetPosts(c *gin.Context) {
	// some logic
}

// CreatePost godoc
//
//	@Summary		Create new Post
//	@Description	Create post with all filled fields
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreatePostModel	true	"CreatePostModel object"
//
//	@Success		201		{object}	Post
//
//	@Failure		400		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/posts/ [post]
func CreatePost(_ *gin.Context) {
	// some logic
}

// UpdatePost godoc
//
//	@Summary		Update post by id
//	@Description	Return updated post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			body	body		UpdatePostModel	true	"UpdatePostModel object"
//
//	@Success		202		{object}	Post
//
//	@Failure		400		{object}	RestError
//	@Failure		404		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/posts/{id} [patch]
func UpdatePost(_ *gin.Context) {
	// some logic
}

// DeletePost godoc
//
//	@Summary		Delete post by id
//	@Description	Delete post by specific id
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"post ID"
//
//	@Success		204
//
//	@Failure		400	{object}	RestError
//	@Failure		404	{object}	RestError
//	@Failure		409	{object}	RestError
//	@Failure		500	{object}	RestError
//	@Failure		502	{object}	RestError
//
//	@Router			/posts/{id} [delete]
func DeletePost(_ *gin.Context) {
	// some logic
}

// GetPostComments godoc
//
//	@Summary		Get all post comments
//	@Description	Return comments from post by specific id
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			limit	path		int	true	"limit for page view"
//	@Param			offset	path		int	true	"offset for page view"
//	@Param			id		path		int	true	"post ID"
//
//	@Success		200		{object}	Comment
//
//	@Failure		400		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/posts/{id}/comments [get]
func GetPostComments(_ *gin.Context) {
	// some logic
}

// GetPostRating godoc
//
//	@Summary		Retur post rating
//	@Description	Return rating of specific post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"post ID"
//
//	@Success		200	{object}	Rating
//
//	@Failure		400	{object}	RestError
//	@Failure		422	{object}	RestError
//	@Failure		500	{object}	RestError
//	@Failure		502	{object}	RestError
//
//	@Router			/posts/{id}/rating [get]
func GetPostRating(_ *gin.Context) {
	// some logic
}

// Rating

// SetRating godoc
//
//	@Summary		Create rating
//	@Description	Create rating by it id
//	@Tags			rating
//	@Accept			json
//	@Produce		json
//	@Param			body	body		RatingModel	true	"RatingModel object"
//
//	@Success		201		{object}	Rating
//
//	@Failure		400		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/rating/ [post]
func SetRating(_ *gin.Context) {
	// some logic
}

// UpdateRating godoc
//
//	@Summary		Update rating
//	@Description	Update rating by it id
//	@Tags			rating
//	@Accept			json
//	@Produce		json
//	@Param			body	body		RatingModel	true	"RatingModel object"
//
//	@Success		201		{object}	Rating
//
//	@Failure		400		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/rating/{id} [patch]
func UpdateRating(_ *gin.Context) {
	// some logic
}

// DeleteRating godoc
//
//	@Summary		Delete rating by id
//	@Description	Delete rating by specific id
//	@Tags			rating
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"rating ID"
//
//	@Success		204
//	@Failure		400	{object}	RestError
//	@Failure		404	{object}	RestError
//	@Failure		409	{object}	RestError
//	@Failure		500	{object}	RestError
//	@Failure		502	{object}	RestError
//
//	@Router			/rating/{id} [delete]
func DeleteRating(_ *gin.Context) {
	// some logic
}

// Comment

// CreateComment godoc
//
//	@Summary		Create comment
//	@Description	Create comment by it id
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreateCommentModel	true	"CreateCommentModel object"
//
//	@Success		201		{object}	Comment
//
//	@Failure		400		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/comments/{id} [post]
func CreateComment(_ *gin.Context) {
	// some logic
}

// UpdateComment godoc
//
//	@Summary		Update comment
//	@Description	Update comment by it id
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			body	body		UpdateCommentModel	true	"UpdateCommentModel object"
//
//	@Success		201		{object}	Comment
//
//	@Failure		400		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/comments/{id} [patch]
func UpdateComment(_ *gin.Context) {
	// some logic
}

// DeleteComment godoc
//
//	@Summary		Delete comment by id
//	@Description	Delete comment by specific id
//	@Tags			comments
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"comment ID"
//
//	@Success		204
//	@Failure		400	{object}	RestError
//	@Failure		404	{object}	RestError
//	@Failure		409	{object}	RestError
//	@Failure		500	{object}	RestError
//	@Failure		502	{object}	RestError
//
//	@Router			/comments/{id} [delete]
func DeleteComment(_ *gin.Context) {
	// some logic
}

// Locations

// CreateLocations godoc
//
//	@Summary		Create location
//	@Description	Create new location
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			body	body		LocationModel	true	"LocationModel object"
//
//	@Success		201		{object}	Location
//
//	@Failure		400		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/locations/{id} [post]
func CreateLocations(_ *gin.Context) {
	// some logic
}

// UpdateLocations godoc
//
//	@Summary		Update location
//	@Description	Update location by it id
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			body	body		LocationModel	true	"LocationModel object"
//
//	@Success		201		{object}	Location
//
//	@Failure		400		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/locations/{id} [patch]
func UpdateLocations(_ *gin.Context) {
	// some logic
}

// DeleteLocations godoc
//
//	@Summary		Delete location by id
//	@Description	Delete location by specific id
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"location ID"
//
//	@Success		204
//	@Failure		400	{object}	RestError
//	@Failure		404	{object}	RestError
//	@Failure		409	{object}	RestError
//	@Failure		500	{object}	RestError
//	@Failure		502	{object}	RestError
//
//	@Router			/locations/{id} [delete]
func DeleteLocations(_ *gin.Context) {
	// some logic
}

// SearchByLocation godoc
//
//	@Summary		Search posts By Location name
//	@Description	Search posts By Location name
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			limit	path		int				true	"limit for page view"
//	@Param			offset	path		int				true	"offset for page view"
//	@Param			body	body		LocationModel	true	"LocationModel object"
//
//	@Success		201		{object}	Comment
//
//	@Failure		400		{object}	RestError
//	@Failure		409		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//
//	@Router			/search/filter-fields [get]
func SearchByLocation(_ *gin.Context) {
	// some logic
}

// GetAllLocations godoc
//
//	@Summary		Get all locations
//	@Description	Return all locations
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			limit	path		int	true	"limit for page view"
//	@Param			offset	path		int	true	"offset for page view"
//	@Success		200		{array}		Location
//	@Failure		400		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//	@Router			/locations/ [get]
func GetAllLocations(ctx *gin.Context) {
	// some logic
}

// GetPostsByLocation godoc
//
//	@Summary		Get all posts by locations
//	@Description	Get all posts by locations
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			limit	path		int	true	"limit for page view"
//	@Param			offset	path		int	true	"offset for page view"
//	@Param			id		path		int	true	"location ID"
//	@Success		200		{array}		Post
//	@Failure		400		{object}	RestError
//	@Failure		422		{object}	RestError
//	@Failure		500		{object}	RestError
//	@Failure		502		{object}	RestError
//	@Router			/locations/{id}/posts [get]
func GetPostsByLocation(_ *gin.Context) {
	// some logic
}
