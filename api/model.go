package main

import (
	"time"

	types "github.com/google/uuid"
)

type (
	Post struct {
		ID           types.UUID `json:"id"`
		Text         string     `json:"text"`
		UserID       types.UUID `json:"user_id"`
		LocationID   types.UUID `json:"location_id"`
		Created      time.Time  `json:"created"`
		LastModified string     `json:"last_modified"`
		Comments     []Comment  `json:"comments,omitempty"`
		Media        []Media    `json:"media,omitempty"`
	}

	Media struct {
		Type  string `json:"type"`
		Photo []byte `json:"photo"`
	}

	Comment struct {
		ID           types.UUID `json:"id"`
		UserID       types.UUID `json:"user_id"`
		Text         string     `json:"text"`
		Created      time.Time  `json:"created"`
		LastModified string     `json:"last_modified"`
	}

	Location struct {
		ID   types.UUID `json:"id"`
		Name string     `json:"name" validate:"required" example:"Hamunaptra,Egypt"`
	}

	Rating struct {
		Points uint8 `json:"point" validate:"required"`
	}

	Subscribe struct {
		ID       types.UUID `json:"id"`
		AuthorID types.UUID `json:"author_id" validate:"required"`
		UserID   types.UUID `json:"user_id" validate:"required"`
	}

	RestError struct {
		Type            string `json:"error_type"`
		Message         string `json:"message" example:"Object with id=%s doesn't exist"`
		AddititonalInfo string `json:"additional_info" example:"127"`
	}
)

type (
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
	Search struct {
		Location string `json:"location"`
	}
)
