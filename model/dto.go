package model

import "time"

type Note struct {
	ID         string `json:"id" form:"id" tag:"id"`
	Content    string `json:"content" form:"content" tag:"content"`
	StartTime  time.Time
	UpdateTime time.Time
}

type NewNote struct {
	ID         int64     `db:"id" form:"id" json:"ID"`
	Content    string    `db:"content" form:"content" json:"Content"`
	StartTime  time.Time `db:"start_time" form:"start_time" json:"StartTime"`
	UpdateTime time.Time `db:"update_time" form:"update_time" json:"UpdateTime"`
}

type NoteRequest struct {
	Offset int `json:"offset" form:"offset" validate:"gte=0"`
	Size   int `json:"size" form:"size" validate:"gte=0"`
}

type EmptyRequest struct{}

type EmptyResponse struct{}
