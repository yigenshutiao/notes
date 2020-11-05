package model

import "time"

type Note struct {
	ID         string `json:"id" form:"id" validate:"required"`
	Content    string `json:"content" form:"content" validate:"required"`
	StartTime  time.Time
	UpdateTime time.Time
}

type EmptyRequest struct{}

type EmptyResponse struct{}
