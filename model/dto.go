package model

import "time"

type Note struct {
	ID         string `json:"id" form:"id" tag:"id"`
	Content    string `json:"content" form:"content" tag:"content"`
	StartTime  time.Time
	UpdateTime time.Time
}

type NewNote struct {
	ID         int64     `db:"id" form:"id"`
	Content    string    `db:"content" form:"content"`
	StartTime  time.Time `db:"start_time" form:"start_time"`
	UpdateTime time.Time `db:"update_time" form:"update_time"`
}

type EmptyRequest struct{}

type EmptyResponse struct{}
