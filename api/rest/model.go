package rest

import (
	"time"

	types "github.com/google/uuid"
)

type (
	User struct {
		ID       types.UUID `json:"id"`
		Name     string     `json:"name" validate:"required"`
		Email    string     `json:"email" validate:"required,email"`
		Password []byte     `json:"password" validate:"required,password"`

		Created      time.Time `json:"created"`
		LastModified string    `json:"last_modified"`
	}

	Subscription struct {
		AuthorID types.UUID `json:"author_id"`
		UserID   types.UUID `json:"user_id"`

		Created      time.Time `json:"created"`
		LastModified string    `json:"last_modified"`
	}

	Post struct {
		ID          types.UUID `json:"id"`
		Description string     `json:"description"`
		UserID      types.UUID `json:"user_id"`
		LocationID  types.UUID `json:"location_id"`

		Created      time.Time `json:"created"`
		LastModified string    `json:"last_modified"`

		Comments []Comment `json:"comments,omitempty"`
		Media    []Media   `json:"media,omitempty"`
		Rating   Rating    `json:"rating,omitempty"`
	}

	Media struct {
		ID types.UUID `json:"id"`

		Type  string `json:"type"`
		Name  string `json:"file_name"`
		Photo []byte `json:"photo"`

		Created      time.Time `json:"created"`
		LastModified string    `json:"last_modified"`
	}

	Comment struct {
		ID     types.UUID `json:"id"`
		PostID types.UUID `json:"post_id"`
		UserID types.UUID `json:"user_id"`
		Text   string     `json:"text"`

		Created      time.Time `json:"created"`
		LastModified string    `json:"last_modified"`
	}

	Location struct {
		ID   types.UUID `json:"id"`
		Name string     `json:"name" validate:"required" example:"Hamunaptra,Egypt"`

		Created      time.Time `json:"created"`
		LastModified string    `json:"last_modified"`
	}

	Rating struct {
		Points uint8 `json:"point" validate:"required"`

		Created      time.Time `json:"created"`
		LastModified string    `json:"last_modified"`
	}

	Subscribe struct {
		AuthorID     types.UUID `json:"author_id" validate:"required"`
		UserID       types.UUID `json:"user_id" validate:"required"`
		Created      time.Time  `json:"created"`
		LastModified string     `json:"last_modified"`
	}

	RestError struct {
		Type            string `json:"error_type"`
		Message         string `json:"message" example:"Object with id=%s doesn't exist"`
		AddititonalInfo string `json:"additional_info" example:"127"`
	}
)

type (
	CreateUserModel struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password []byte `json:"password" validate:"required,password"`
	}

	UpdateUserModel struct {
		Name         string `json:"name"`
		Email        string `json:"email"`
		Password     []byte `json:"password"`
		LastModified string `json:"last_modified" validate:"required"`
	}

	SubscriptionModel struct {
		AuthorID types.UUID `json:"author_id"`
		UserID   types.UUID `json:"user_id"`
	}
	CreatePostModel struct {
		Text       string     `json:"text"`
		LocationID types.UUID `json:"location_id"`
		Media      []Media    `json:"media,omitempty"`
	}

	UpdatePostModel struct {
		Text         string     `json:"text"`
		LocationID   types.UUID `json:"location_id"`
		Media        []Media    `json:"media,omitempty"`
		LastModified string     `json:"last_modified" validate:"required"`
	}

	CreateCommentModel struct {
		Text   string     `json:"text"`
		PostID types.UUID `json:"post_id"`
	}

	UpdateCommentModel struct {
		Text         string     `json:"text"`
		PostID       types.UUID `json:"post_id"`
		LastModified string     `json:"last_modified"`
	}

	RatingModel struct {
		Points uint8 `json:"point" validate:"required"`
	}

	LocationModel struct {
		Name string `json:"name"`
	}
)
