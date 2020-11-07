package model

import "time"

type Note struct {
	ID         string `json:"id" form:"id" tag:"id"`
	Content    string `json:"content" form:"content" tag:"content"`
	StartTime  time.Time
	UpdateTime time.Time
}

type EmptyRequest struct{}

type EmptyResponse struct{}
